package showonce_test

import (
	"testing"

	app "github.com/LeKovr/showonce"
	gen "github.com/LeKovr/showonce/zgen/go/proto"
	"github.com/stretchr/testify/assert"
)

func TestFlow(t *testing.T) {
	cfg := app.StorageConfig{}
	db := app.NewStorage(cfg)

	item := &gen.NewItemRequest{
		Title:      "title",
		Group:      "group",
		Data:       "data",
		Expire:     "1",
		ExpireUnit: "d",
	}
	user := "test"

	id, err := db.SetMeta(user, item)
	assert.NoError(t, err, "SetMeta")
	assert.NotNil(t, id, "SetMetaNotNil")

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

	_, err = db.GetData(id.String())
	assert.ErrorIs(t, err, app.ErrNotFound, "GetData2")

	/*
	   sleep > dataTTL => no data
	   sleep > metaTTL => no meta
	   cleanup interval?
	*/
}

/*
tests := []struct {
		isOk    bool
		version string
		repo    string
		err     string
	}{
		{true, "v0.31", "https://github.com/LeKovr/dbrpc.git", ""},
		{true, "v0.31", "git@github.com:LeKovr/dbrpc.git", ""},
		{true, "any version is ok", "git@github.com:LeKovr/golang-use.git", ""},
		{false, "v0.30", "https://github.com/LeKovr/dbrpc.git", "{\"level\":\"info\",\"v\":0,\"appVersion\":\"v0.30\",\"sourceVersion\":\"v0.31\",\"sourceUpdated\":\"2017-10-17T08:56:03Z\",\"sourceLink\":\" See https://github.com/LeKovr/dbrpc/releases/tag/v0.31\",\"message\":\"App version is outdated\"}\n"},
		{false, "v0.0", "https://localhost:10", "Get \"https://localhost:10/releases.atom\": dial tcp 127.0.0.1:10: connect: connection refused"},
	}
	for _, tt := range tests {
		buf := new(bytes.Buffer)
		zl := zerolog.New(buf).Level(zerolog.InfoLevel)
		var log logr.Logger = zerologr.New(&zl)
		ok, err := ver.IsCheckOk(log, tt.repo, tt.version)
		assert.Equal(t, tt.isOk, ok)
		if !tt.isOk {
			if err != nil {
				assert.EqualError(t, err, tt.err)
			} else {
				assert.Equal(t, tt.err, buf.String())
			}
		}
	}
}
*/
