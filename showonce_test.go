package showonce_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	app "github.com/LeKovr/showonce"
	storage "github.com/LeKovr/showonce/storage/cache"
	gen "github.com/LeKovr/showonce/zgen/go/proto"
	test_suite "github.com/stretchr/testify/suite"
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
	ass := suite.Require()
	ass.NoError(err, "NewItem")
	ass.NotNil(id, "NewItem returns not nil")

	meta, err := suite.Pub.GetMetadata(ctx, id)
	ass.NoError(err, "GetMeta")
	ass.Equal(item.GetGroup(), meta.GetGroup(), "GetMetaGroupEq")
	ass.Equal(suite.User, meta.GetOwner(), "GetMetaOwnerEq")

	stats, err := suite.Priv.GetStats(ctxMD, empty)
	ass.NoError(err, "GetStats")
	ass.Equal(int32(1), stats.GetMy().GetTotal(), "My Items Total must be 1")

	items, err := suite.Priv.GetItems(ctxMD, empty)
	ass.NoError(err, "GetItems")
	ass.Equal(gen.ItemStatus_WAIT, items.GetItems()[0].GetMeta().GetStatus(), "ItemStatusIsWait")

	data, err := suite.Pub.GetData(ctx, id)
	ass.NoError(err, "GetData")
	ass.Equal(item.GetData(), data.GetData(), "GetDataEq")

	_, err = suite.Pub.GetData(ctx, id)
	e, ok := status.FromError(err)
	ass.True(ok, "Error is GRPC error")
	ass.Equal(codes.NotFound, e.Code(), "GetDataIsEmpty")

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
	ass := suite.Require()

	id, err := suite.Priv.NewItem(ctxMD, item)
	ass.NoError(err, "NewItem")
	ass.NotNil(id, "NewItem returns not nil")

	exp, _ := time.ParseDuration(item.GetExpire() + item.GetExpireUnit())
	time.Sleep(exp)
	meta, err := suite.Pub.GetMetadata(ctx, id)
	ass.NoError(err, "GetItems")
	ass.Equal(gen.ItemStatus_EXPIRED, meta.GetStatus(), "ItemStatusIsExpired")
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
	ass := suite.Require()
	for _, tt := range tests {
		ctx := context.Background()
		if tt.md != nil {
			ctx = metadata.NewIncomingContext(ctx, tt.md)
		}
		_, err := suite.Priv.GetStats(ctx, empty)
		e, ok := status.FromError(err)
		ass.True(ok, "Error is GRPC error")
		ass.Equal(codes.PermissionDenied, e.Code(), tt.name)
	}
}
