package showonce_test

import (
	"context"
	"testing"
	"time"

	app "github.com/LeKovr/showonce"
	storage "github.com/LeKovr/showonce/storage/cache"
	gen "github.com/LeKovr/showonce/zgen/go/proto"
	test_suite "github.com/stretchr/testify/suite"
	"google.golang.org/grpc/metadata"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type ShowonceTestSuite struct {
	test_suite.Suite
	User string
	Pub  *app.PublicServiceImpl
	Priv *app.PrivateServiceImpl
}

var empty = &emptypb.Empty{}

func (suite *ShowonceTestSuite) SetupTest() {
	cfg := storage.Config{}
	db := storage.New(cfg)

	suite.User = "test"
	suite.Pub = app.NewPublicService(db)
	suite.Priv = app.NewPrivateService(db)
}

func TestShowonceTestSuite(t *testing.T) {
	test_suite.Run(t, new(ShowonceTestSuite))
}

func (suite *ShowonceTestSuite) TestFlow() {
	item := &gen.NewItemRequest{
		Title:      "title",
		Group:      "group",
		Data:       "data",
		Expire:     "1",
		ExpireUnit: "d",
	}
	ctx := context.Background()
	ctxMD := metadata.NewIncomingContext(ctx, metadata.Pairs(app.MDUserKey, suite.User))

	id, err := suite.Priv.NewItem(ctxMD, item)
	suite.NoError(err, "NewItem")
	suite.NotNil(id, "NewItem returns not nil")

	meta, err := suite.Pub.GetMetadata(ctx, id)
	suite.NoError(err, "GetMeta")
	suite.Equal(item.Group, meta.Group, "GetMetaGroupEq")
	suite.Equal(suite.User, meta.Owner, "GetMetaOwnerEq")

	stats, err := suite.Priv.GetStats(ctxMD, empty)
	suite.NoError(err, "GetStats")
	suite.Equal(int32(1), stats.My.Total, "My Items Total must be 1")

	items, err := suite.Priv.GetItems(ctxMD, empty)
	suite.NoError(err, "GetItems")
	suite.Equal(gen.ItemStatus_WAIT, items.Items[0].Meta.Status, "ItemStatusIsWait")

	data, err := suite.Pub.GetData(ctx, id)
	suite.NoError(err, "GetData")
	suite.Equal(item.Data, data.Data, "GetDataEq")

	_, err = suite.Pub.GetData(ctx, id)
	suite.ErrorIs(err, storage.ErrNotFound, "GetDataIsEmpty")

	/*
	   sleep > dataTTL => no data
	   sleep > metaTTL => no meta
	   cleanup interval?
	*/
}

func (suite *ShowonceTestSuite) TestExpire() {
	item := &gen.NewItemRequest{
		Title:      "title",
		Group:      "group",
		Data:       "data",
		Expire:     "10",
		ExpireUnit: "ms",
	}
	ctx := context.Background()
	ctxMD := metadata.NewIncomingContext(ctx, metadata.Pairs(app.MDUserKey, suite.User))

	id, err := suite.Priv.NewItem(ctxMD, item)
	suite.NoError(err, "NewItem")
	suite.NotNil(id, "NewItem returns not nil")

	exp, _ := time.ParseDuration(item.Expire + item.ExpireUnit)
	time.Sleep(exp)
	meta, err := suite.Pub.GetMetadata(ctx, id)
	suite.NoError(err, "GetItems")
	suite.Equal(gen.ItemStatus_EXPIRED, meta.Status, "ItemStatusIsExpired")
}

func (suite *ShowonceTestSuite) TestAuthErrors() {
	tests := []struct {
		name string
		md   metadata.MD
	}{
		{"no metadata", nil},
		{"no field 'user'", metadata.Pairs("UNKNOWN", suite.User)},
		{"field 'user' is empty", metadata.Pairs(app.MDUserKey, "")},
	}
	for _, tt := range tests {
		ctx := context.Background()
		if tt.md != nil {
			ctx = metadata.NewIncomingContext(ctx, tt.md)
		}
		_, err := suite.Priv.GetStats(ctx, empty)
		suite.ErrorIs(err, app.ErrMetadataMissing, tt.name)
	}
}
