/*
Package showonce - реализация публичного и приватного сервиса.
*/
package showonce

import (
	"context"

	"github.com/LeKovr/showonce/storage"
	gen "github.com/LeKovr/showonce/zgen/go/proto"
	"github.com/go-logr/logr"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

var errMissingMetadata = status.Errorf(codes.InvalidArgument, "no incoming metadata in rpc context")

// PublicServiceImpl - реализация PublicService.
type PublicServiceImpl struct {
	gen.UnimplementedPublicServiceServer
	Store storage.StorageIface
}

// NewPublicService - создать PublicService.
func NewPublicService(db storage.StorageIface) *PublicServiceImpl {
	return &PublicServiceImpl{Store: db}
}

// GetMetadata - вернуть метаданные по id.
func (service PublicServiceImpl) GetMetadata(_ context.Context, id *gen.ItemId) (*gen.ItemMeta, error) {
	rv, err := service.Store.GetMeta(id.Id)
	return rv, err
}

// GetData -вернуть контент по id.
func (service PublicServiceImpl) GetData(_ context.Context, id *gen.ItemId) (*gen.ItemData, error) {
	rv, err := service.Store.GetData(id.Id)
	return rv, err
}

// PrivateServiceImpl - реадизация PrivateService.
type PrivateServiceImpl struct {
	gen.UnimplementedPrivateServiceServer
	Store storage.StorageIface
}

// NewPrivateService - создать PrivateService.
func NewPrivateService(db storage.StorageIface) *PrivateServiceImpl {
	return &PrivateServiceImpl{Store: db}
}

// NewMessage - создать контент.
func (service PrivateServiceImpl) NewMessage(ctx context.Context, req *gen.NewItemRequest) (*gen.ItemId, error) {
	user, err := fetchUser(ctx)
	if err != nil {
		return nil, err
	}

	idStr, err := service.Store.SetMeta(*user, req)
	if err != nil {
		log := logr.FromContextOrDiscard(ctx)
		log.Error(err, "NewMessageError")
		return nil, err
	}
	return &gen.ItemId{Id: idStr.String()}, nil
}

// GetItems - вернуть список своих текстов.
func (service PrivateServiceImpl) GetItems(ctx context.Context, _ *emptypb.Empty) (*gen.ItemList, error) {
	user, err := fetchUser(ctx)
	if err != nil {
		return nil, err
	}
	rv, err := service.Store.Items(*user)
	return rv, err
}

// GetStats - общая статистика (всего/активных текстов, макс дата активного текста).
func (service PrivateServiceImpl) GetStats(ctx context.Context, _ *emptypb.Empty) (*gen.StatsResponse, error) {
	log := logr.FromContextOrDiscard(ctx)
	log.Info("GetStats")
	user, err := fetchUser(ctx)
	if err != nil {
		return nil, err
	}
	rv, err := service.Store.Stats(*user)
	return rv, err
}

func fetchUser(ctx context.Context) (*string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}
	// Fetch Username
	user := md["user"][0]
	log := logr.FromContextOrDiscard(ctx)
	if user == "" {
		log.Info("Username must be set")
		return nil, errMissingMetadata
	}
	log.Info("USER", "name", user)
	return &user, nil
}
