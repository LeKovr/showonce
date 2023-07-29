package showonce_test

import (
	"context"
	"testing"

	app "github.com/LeKovr/showonce"
	storage "github.com/LeKovr/showonce/storage/cache"
	gen "github.com/LeKovr/showonce/zgen/go/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestRPC(t *testing.T) {
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
	empty := &emptypb.Empty{}
	ctx := context.Background()
	pub := app.NewPublicService(db)
	priv := app.NewPrivateService(db)
	ctxMD := metadata.NewIncomingContext(ctx, metadata.Pairs("user", user))
	id, err := priv.NewMessage(ctxMD, item)
	assert.NoError(t, err, "SetMeta")
	assert.NotNil(t, id, "SetMetaNotNil")

	meta, err := pub.GetMetadata(ctx, id)
	assert.NoError(t, err, "GetMeta")
	assert.Equal(t, item.Group, meta.Group, "GetMetaEq")
	assert.Equal(t, user, meta.Owner, "GetMetaOwnerEq")

	stats, err := priv.GetStats(ctxMD, empty)
	assert.NoError(t, err, "GetStats")
	assert.Equal(t, int32(1), stats.My.Total, "My Items Total must be 1")

	items, err := priv.GetItems(ctxMD, empty)
	assert.NoError(t, err, "GetItems")
	assert.Equal(t, gen.ItemStatus_WAIT, items.Items[0].Meta.Status, "ItemsStatusIsWait")

	data, err := pub.GetData(ctx, id)
	assert.NoError(t, err, "GetData")
	assert.Equal(t, item.Data, data.Data, "GetDataEq")

	_, err = pub.GetData(ctx, id)
	assert.ErrorIs(t, err, storage.ErrNotFound, "GetData2")

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
