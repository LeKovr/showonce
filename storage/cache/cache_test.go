package cache_test

import (
	"testing"

	storerr "github.com/LeKovr/showonce/storage"
	storage "github.com/LeKovr/showonce/storage/cache"
	gen "github.com/LeKovr/showonce/zgen/go/proto"
	ass "github.com/alecthomas/assert/v2"
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
	ass.NoError(t, err, "SetItem")
	ass.NotZero(t, id, "SetItemNotNil")

	meta, err := db.GetMeta(id.String())
	ass.NoError(t, err, "GetMeta")
	ass.Equal(t, item.GetGroup(), meta.GetGroup(), "GetMetaEq")
	ass.Equal(t, user, meta.GetOwner(), "GetMetaOwnerEq")

	stats, err := db.Stats(user)
	ass.NoError(t, err, "Stats")
	ass.Equal(t, int32(1), stats.GetMy().GetTotal(), "My Items Total must be 1")

	items, err := db.Items(user)
	ass.NoError(t, err, "Items")
	ass.Equal(t, gen.ItemStatus_WAIT, items.GetItems()[0].GetMeta().GetStatus(), "ItemsStatusIsWait")

	data, err := db.GetData(id.String())
	ass.NoError(t, err, "GetData")
	ass.Equal(t, data.GetData(), item.GetData(), "GetDataEq")

	meta, err = db.GetMeta(id.String())
	ass.NoError(t, err, "GetMeta")
	ass.Equal(t, meta.GetStatus(), gen.ItemStatus_READ, "GetMetaEqRead")

	_, err = db.GetData(id.String())
	ass.IsError(t, err, storerr.ErrNotFound, "GetData2")

	/*
	   sleep > dataTTL => no data
	   sleep > metaTTL => no meta
	   cleanup interval?
	*/
}
