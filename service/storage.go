package service

import (
	crand "crypto/rand"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/patrickmn/go-cache"
)

type StorageConfig struct {
	MetaTTL         time.Duration `long:"meta_ttl" default:"240h" description:"Metadata TTL"`
	DataTTL         time.Duration `long:"data_ttl" default:"24h" description:"Data TTL"`
	CleanupInterval time.Duration `long:"cleanup" default:"10m" description:"Cleanup interval"`
}

type Storage struct {
	Meta  *cache.Cache
	Owner *cache.Cache
	Data  *cache.Cache
}

var (
	ErrNoUniqueWithinLimit = errors.New("Cannot create unique id")
	ErrNotFound            = errors.New("Item not found")
	ErrDataCorrupted       = errors.New("Item data was corrupted")
)

func NewStorage(cfg StorageConfig) Storage {
	return Storage{
		Meta:  cache.New(cfg.MetaTTL, cfg.CleanupInterval),
		Owner: cache.New(cfg.MetaTTL, cfg.CleanupInterval),
		Data:  cache.New(cfg.DataTTL, cfg.CleanupInterval),
	}
	//TODO - Data. OnEvicted - update Meta.Status
}

func (store Storage) SetMeta(owner string, req NewItemRequest) (*ulid.ULID, error) {

	// Validate expiration
	var expire time.Duration
	if req.ExpireValue != "" {
		if req.ExpireUnit == "d" {
			days, err := strconv.Atoi(req.ExpireValue)
			if err != nil {
				return nil, fmt.Errorf("Expire days parse error: %w", err)
			}
			expire = time.Duration(days) * time.Hour * 24
		} else {
			var err error
			expire, err = time.ParseDuration(fmt.Sprintf("%s%s", req.ExpireValue, req.ExpireUnit))
			if err != nil {
				return nil, fmt.Errorf("Expire parse error: %w", err)
			}
		}
	}
	meta := ItemMeta{
		Item: Item{Title: req.Item.Title,
			Group: req.Group,
		},
		Owner:    owner,
		Created:  time.Now(),
		Modified: time.Now().Add(expire),
		Status:   StatusWait,
	}
	// TODO: meta exp = data exp + meta config

	for i := 0; i < 5; i++ {
		// try to get unique id
		ms := ulid.Timestamp(time.Now())
		id, err := ulid.New(ms, crand.Reader)
		if err != nil {
			return nil, fmt.Errorf("ID generate error: %w", err)
		}

		err = store.Data.Add(id.String(), req.Data, expire)
		if err == nil {
			// data is unique
			store.Meta.SetDefault(id.String(), &meta)
			return &id, nil

		}
	}
	return nil, ErrNoUniqueWithinLimit

}

func (store Storage) GetMeta(id string) (*ItemMeta, error) {
	meta, ok := store.Meta.Get(id)
	if !ok {
		return nil, ErrNotFound
	}
	rv, ok := meta.(*ItemMeta)
	if !ok {
		return nil, ErrDataCorrupted
	}
	checkExpire(rv)
	return rv, nil
}

func (store Storage) GetData(id string) (*string, error) {
	data, ok := store.Data.Get(id)
	if !ok {
		return nil, ErrNotFound
	}
	rv, ok := data.(string)
	if !ok {
		return nil, ErrDataCorrupted
	}
	meta, err := store.GetMeta(id)
	if err != nil {
		return nil, err
	}
	meta.Status = StatusRead
	meta.Modified = time.Now()

	err = store.Meta.Replace(id, meta, 0)
	store.Data.Delete(id)
	return &rv, nil
}

func (store Storage) Items(owner string) (items []ItemInfo, err error) {
	cacheItems := store.Meta.Items()
	for k, v := range cacheItems {
		meta := v.Object.(*ItemMeta)
		if meta.Owner != owner {
			continue
		}
		checkExpire(meta)
		items = append(items, ItemInfo{Id: k, Meta: meta})
	}
	return
}

func (store Storage) Stats(owner string) (stat StatResponse, err error) {
	cacheItems := store.Meta.Items()
	for _, v := range cacheItems {
		meta := v.Object.(*ItemMeta)
		checkExpire(meta)
		var curr *Stat
		if meta.Owner == owner {
			curr = &stat.My
		} else {
			curr = &stat.Other
		}
		curr.Total++
		switch meta.Status {
		case StatusWait:
			curr.Wait++
		case StatusRead:
			curr.Read++
		case StatusExpired:
			curr.Expired++
		}
	}
	return
}

// TODO: replace with OnEvicted
func checkExpire(meta *ItemMeta) {
	if meta.Status == StatusWait && time.Now().After(meta.Modified) {
		// Data expired already
		meta.Status = StatusExpired
	}
}
