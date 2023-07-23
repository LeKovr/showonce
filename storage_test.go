package showonce_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	app "github.com/LeKovr/showonce"
	gen "github.com/LeKovr/showonce/zgen/go/proto"
)

func TestFlow(t *testing.T) {
	cfg := app.StorageConfig{}
	db := app.NewStorage(cfg)

	item := &gen.NewItemRequest{Title: "title", Group: "group", Data: "data"}
	user := "test"

	id, err := db.SetMeta(user, item)
	assert.NoError(t, err, "SetMeta")
	assert.NotNil(t, id, "SetMetaNotNil")

	meta, err := db.GetMeta(id.String())
	assert.NoError(t, err, "GetMeta")
	assert.Equal(t, item.Group, meta.Group, "GetMetaEq")
	assert.Equal(t, user, meta.Owner, "GetMetaOwnerEq")

	data, err := db.GetData(id.String())
	assert.NoError(t, err, "GetData")
	assert.Equal(t, item.Data, data.Data, "GetDataEq")

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
