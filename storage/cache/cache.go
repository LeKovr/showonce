// Package cache implements cache storage.
package cache

import (
	crand "crypto/rand"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"github.com/oklog/ulid/v2"
	"google.golang.org/protobuf/types/known/timestamppb"
	zcache "zgo.at/zcache/v2"

	storerr "github.com/LeKovr/showonce/storage"
	gen "github.com/LeKovr/showonce/zgen/go/proto"
)

// Config holds Storage Config.
type Config struct {
	MetaTTL         time.Duration `default:"240h" description:"Metadata TTL"     long:"meta_ttl"`
	DataTTL         time.Duration `default:"24h"  description:"Data TTL"         long:"data_ttl"`
	CleanupInterval time.Duration `default:"10m"  description:"Cleanup interval" long:"cleanup"`
}

// Storage implements data storage.
type Storage struct {
	Meta    *zcache.Cache[string, *gen.ItemMeta]
	Data    *zcache.Cache[string, string]
	DataTTL time.Duration
}

const hoursInDay = 24

// New returns new Storage object.
func New(cfg Config) Storage {
	meta := zcache.New[string, *gen.ItemMeta](cfg.MetaTTL, cfg.CleanupInterval)
	data := zcache.New[string, string](cfg.DataTTL, cfg.CleanupInterval)
	data.OnEvicted(func(k, _ string) {
		// Set metadata status when data expires
		if item, ok := meta.Get(k); ok {
			if item.GetStatus() == gen.ItemStatus_WAIT {
				item.Status = gen.ItemStatus_EXPIRED
				meta.Set(k, item)
			}
		}
	})

	return Storage{Meta: meta, Data: data, DataTTL: cfg.DataTTL}
}

// SetItem prepares and saves item metadata and secret.
func (store Storage) SetItem(owner string, req *gen.NewItemRequest) (*ulid.ULID, error) {
	// Validate expiration
	var expire time.Duration

	if req.GetExpire() != "" {
		if req.GetExpireUnit() == "d" {
			days, err := strconv.Atoi(req.GetExpire())
			if err != nil {
				return nil, fmt.Errorf("expire days parse error: %w", err)
			}

			expire = time.Duration(days) * time.Hour * hoursInDay
		} else {
			var err error

			expire, err = time.ParseDuration(fmt.Sprintf("%s%s", req.GetExpire(), req.GetExpireUnit()))
			if err != nil {
				return nil, fmt.Errorf("expire parse error: %w", err)
			}
		}
	} else {
		expire = store.DataTTL
	}

	now := time.Now()
	meta := gen.ItemMeta{
		Title:      req.GetTitle(),
		Group:      req.GetGroup(),
		Owner:      owner,
		CreatedAt:  timestamppb.New(now),
		ModifiedAt: timestamppb.New(now.Add(expire)),
		Status:     gen.ItemStatus_WAIT,
	}
	// TODO: meta exp = data exp + meta config

	for range 5 {
		// try to get unique id
		ms := ulid.Timestamp(time.Now())

		id, err := ulid.New(ms, crand.Reader)
		if err != nil {
			return nil, fmt.Errorf("ID generate error: %w", err)
		}

		err = store.Data.AddWithExpire(id.String(), req.GetData(), expire)
		if err == nil {
			// data is unique
			slog.Debug("New item", "exire", expire)

			err = store.Meta.Add(id.String(), &meta)

			return &id, err
		}
	}

	return nil, storerr.ErrNoUniqueWithinLimit
}

// GetMeta returns item metadata.
func (store Storage) GetMeta(id string) (*gen.ItemMeta, error) {
	meta, ok := store.Meta.Get(id)
	if !ok {
		return nil, storerr.ErrNotFound
	}

	checkExpire(meta)

	return meta, nil
}

// GetData returns item data (secret).
func (store Storage) GetData(id string) (*gen.ItemData, error) {
	data, ok := store.Data.Get(id)
	if !ok {
		return nil, storerr.ErrNotFound
	}

	meta, err := store.GetMeta(id)
	if err != nil {
		return nil, err
	}

	meta.Status = gen.ItemStatus_READ
	meta.ModifiedAt = timestamppb.Now()

	err = store.Meta.Replace(id, meta)
	if err != nil {
		return nil, err
	}

	store.Data.Delete(id)
	rv := gen.ItemData{Data: data}

	return &rv, nil
}

// Items returns items, created by current user.
func (store Storage) Items(owner string) (*gen.ItemList, error) {
	cacheItems := store.Meta.Items()
	items := &gen.ItemList{Items: []*gen.ItemMetaWithId{}}

	for k, v := range cacheItems {
		meta := v.Object
		if meta.GetOwner() != owner {
			continue
		}

		checkExpire(meta)
		items.Items = append(items.GetItems(), &gen.ItemMetaWithId{Id: k, Meta: meta})
	}

	return items, nil
}

// Stats returns global and user's item counters.
func (store Storage) Stats(owner string) (*gen.StatsResponse, error) {
	cacheItems := store.Meta.Items()
	stat := &gen.StatsResponse{My: &gen.Stats{}, Other: &gen.Stats{}}

	for _, v := range cacheItems {
		meta := v.Object
		checkExpire(meta)

		var curr *gen.Stats
		if meta.GetOwner() == owner {
			curr = stat.GetMy()
		} else {
			curr = stat.GetOther()
		}

		curr.Total++

		switch meta.GetStatus() {
		case gen.ItemStatus_WAIT:
			curr.Wait++
		case gen.ItemStatus_READ:
			curr.Read++
		case gen.ItemStatus_EXPIRED:
			curr.Expired++
		case gen.ItemStatus_CLEARED:
			// Ignore
		case gen.ItemStatus_UNKNOWN:
			slog.Warn("Unknown status", "meta", meta)
		}
	}

	return stat, nil
}

// TODO: replace with OnEvicted.
func checkExpire(meta *gen.ItemMeta) {
	if meta.GetStatus() == gen.ItemStatus_WAIT && time.Now().After(meta.GetModifiedAt().AsTime()) {
		// Data expired already
		meta.Status = gen.ItemStatus_EXPIRED
	}
}
