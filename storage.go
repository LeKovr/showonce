package showonce

import (
	crand "crypto/rand"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/oklog/ulid/v2"
	//"github.com/patrickmn/go-cache"
	"google.golang.org/protobuf/types/known/timestamppb"
	cache "zgo.at/zcache/v2"

	gen "github.com/LeKovr/showonce/zgen/go/proto"
)

type StorageConfig struct {
	MetaTTL         time.Duration `long:"meta_ttl" default:"240h" description:"Metadata TTL"`
	DataTTL         time.Duration `long:"data_ttl" default:"24h" description:"Data TTL"`
	CleanupInterval time.Duration `long:"cleanup" default:"10m" description:"Cleanup interval"`
}

type Storage struct {
	Meta *cache.Cache[string, *gen.ItemMeta]
	//Owner *cache.Cache
	Data *cache.Cache[string, string]
}

var (
	ErrNoUniqueWithinLimit = errors.New("Cannot create unique id")
	ErrNotFound            = errors.New("Item not found")
	ErrDataCorrupted       = errors.New("Item data was corrupted")
)

func NewStorage(cfg StorageConfig) Storage {
	return Storage{
		Meta: cache.New[string, *gen.ItemMeta](cfg.MetaTTL, cfg.CleanupInterval),
		//	Owner: cache.New(cfg.MetaTTL, cfg.CleanupInterval),
		Data: cache.New[string, string](cfg.DataTTL, cfg.CleanupInterval),
	}
	//TODO - Data. OnEvicted - update Meta.Status
}

func (store Storage) SetMeta(owner string, req *gen.NewItemRequest) (*ulid.ULID, error) {

	// Validate expiration
	var expire time.Duration
	if req.Expire != "" {
		if req.ExpireUnit == "d" {
			days, err := strconv.Atoi(req.Expire)
			if err != nil {
				return nil, fmt.Errorf("Expire days parse error: %w", err)
			}
			expire = time.Duration(days) * time.Hour * 24
		} else {
			var err error
			expire, err = time.ParseDuration(fmt.Sprintf("%s%s", req.Expire, req.ExpireUnit))
			if err != nil {
				return nil, fmt.Errorf("Expire parse error: %w", err)
			}
		}
	}
	now := time.Now()
	meta := gen.ItemMeta{
		Title:      req.Title,
		Group:      req.Group,
		Owner:      owner,
		CreatedAt:  timestamppb.New(now),
		ModifiedAt: timestamppb.New(now.Add(expire)),
		Status:     gen.ItemStatus_WAIT,
	}
	// TODO: meta exp = data exp + meta config

	for i := 0; i < 5; i++ {
		// try to get unique id
		ms := ulid.Timestamp(time.Now())
		id, err := ulid.New(ms, crand.Reader)
		if err != nil {
			return nil, fmt.Errorf("ID generate error: %w", err)
		}

		err = store.Data.AddWithExpire(id.String(), req.Data, expire)
		if err == nil {
			// data is unique
			store.Meta.Add(id.String(), &meta)
			return &id, nil

		}
	}
	return nil, ErrNoUniqueWithinLimit
}

func (store Storage) GetMeta(id string) (*gen.ItemMeta, error) {
	meta, ok := store.Meta.Get(id)
	if !ok {
		return nil, ErrNotFound
	}
	checkExpire(meta)
	return meta, nil
}

func (store Storage) GetData(id string) (*gen.ItemData, error) {
	data, ok := store.Data.Get(id)
	if !ok {
		return nil, ErrNotFound
	}
	meta, err := store.GetMeta(id)
	if err != nil {
		return nil, err
	}
	meta.Status = gen.ItemStatus_READ
	meta.ModifiedAt = timestamppb.Now()

	err = store.Meta.Replace(id, meta)
	store.Data.Delete(id)
	rv := gen.ItemData{Data: data}
	return &rv, nil
}

func (store Storage) Items(owner string) (items *gen.ItemList, err error) {
	cacheItems := store.Meta.Items()
	items = &gen.ItemList{Items: []*gen.ItemMetaWithId{}}
	for k, v := range cacheItems {
		meta := v.Object
		if meta.Owner != owner {
			continue
		}
		checkExpire(meta)
		items.Items = append(items.Items, &gen.ItemMetaWithId{Id: k, Meta: meta})
	}
	return
}

func (store Storage) Stats(owner string) (stat *gen.StatsResponse, err error) {
	cacheItems := store.Meta.Items()
	stat = &gen.StatsResponse{My: &gen.Stats{}, Other: &gen.Stats{}}
	for _, v := range cacheItems {
		meta := v.Object
		checkExpire(meta)
		var curr *gen.Stats
		if meta.Owner == owner {
			curr = stat.My
		} else {
			curr = stat.Other
		}
		curr.Total++
		switch meta.Status {
		case gen.ItemStatus_WAIT:
			curr.Wait++
		case gen.ItemStatus_READ:
			curr.Read++
		case gen.ItemStatus_EXPIRED:
			curr.Expired++
		}
	}
	return
}

// TODO: replace with OnEvicted
func checkExpire(meta *gen.ItemMeta) {
	if meta.Status == gen.ItemStatus_WAIT && time.Now().After(meta.ModifiedAt.AsTime()) {
		// Data expired already
		meta.Status = gen.ItemStatus_EXPIRED
	}
}
