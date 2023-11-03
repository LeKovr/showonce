package cache_test

import (
	"testing"

	storage "github.com/LeKovr/showonce/storage/cache"
	gen "github.com/LeKovr/showonce/zgen/go/proto"
	"github.com/stretchr/testify/assert"
)

func TestFlow(t *testing.T) {
	cfg := storage.Config{}
	db := storage.New(cfg)

	item := &gen.NewItemRequest{
		Title:      "title",
		Group:      "group",
		Data:       "data",
		Expire:     "1",
		ExpireUnit: "d",
	}
	user := "test"

	id, err := db.SetItem(user, item)
	assert.NoError(t, err, "SetItem")
	assert.NotNil(t, id, "SetItemNotNil")

	meta, err := db.GetMeta(id.String())
	assert.NoError(t, err, "GetMeta")
	assert.Equal(t, item.Group, meta.Group, "GetMetaEq")
	assert.Equal(t, user, meta.Owner, "GetMetaOwnerEq")

	stats, err := db.Stats(user)
	assert.NoError(t, err, "Stats")
	assert.Equal(t, int32(1), stats.My.Total, "My Items Total must be 1")

	items, err := db.Items(user)
	assert.NoError(t, err, "Items")
	assert.Equal(t, gen.ItemStatus_WAIT, items.Items[0].Meta.Status, "ItemsStatusIsWait")

	data, err := db.GetData(id.String())
	assert.NoError(t, err, "GetData")
	assert.Equal(t, data.Data, item.Data, "GetDataEq")

	meta, err = db.GetMeta(id.String())
	assert.NoError(t, err, "GetMeta")
	assert.Equal(t, meta.Status, gen.ItemStatus_READ,  "GetMetaEqRead")

	_, err = db.GetData(id.String())
	assert.ErrorIs(t, err, storage.ErrNotFound, "GetData2")

	/*
	   sleep > dataTTL => no data
	   sleep > metaTTL => no meta
	   cleanup interval?
	*/
}
