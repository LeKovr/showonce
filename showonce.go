/*
Package showonce - реализация публичного и приватного сервиса.
*/
package showonce

import (
	"context"

	gen "github.com/LeKovr/showonce/zgen/go/proto"
	"github.com/go-logr/logr"
	"github.com/oklog/ulid/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// StorageIface makes users independent from storage implementation.
type StorageIface interface {
	// SetItem prepares and saves item metadata and secret.
	SetItem(owner string, req *gen.NewItemRequest) (*ulid.ULID, error)
	// GetMeta returns item metadata.
	GetMeta(id string) (*gen.ItemMeta, error)
	// GetData returns item data (secret).
	GetData(id string) (*gen.ItemData, error)
	// Items returns items, created by current user.
	Items(owner string) (*gen.ItemList, error)
	// Stats returns global and user's item counters.
	Stats(owner string) (*gen.StatsResponse, error)
}

// MDUserKey holds GRPC metadata key for user name.
const MDUserKey = "user"

// ErrMetadataMissing means no user data found in request context.
var ErrMetadataMissing = status.Errorf(codes.InvalidArgument, "no required metadata in rpc context")

// PublicServiceImpl - реализация PublicService.
type PublicServiceImpl struct {
	gen.UnimplementedPublicServiceServer
	Store StorageIface
}

// NewPublicService - создать PublicService.
func NewPublicService(db StorageIface) *PublicServiceImpl {
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
	Store StorageIface
}

// NewPrivateService - создать PrivateService.
func NewPrivateService(db StorageIface) *PrivateServiceImpl {
	return &PrivateServiceImpl{Store: db}
}

// NewItem - создать контент.
func (service PrivateServiceImpl) NewItem(ctx context.Context, req *gen.NewItemRequest) (*gen.ItemId, error) {
	user, err := fetchUser(ctx)
	if err != nil {
		return nil, err
	}

	idStr, err := service.Store.SetItem(*user, req)
	if err != nil {
		log := logr.FromContextOrDiscard(ctx)
		log.Error(err, "NewItemError")
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

// fetchUser fetches user name from ctx metadata.
func fetchUser(ctx context.Context) (*string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, ErrMetadataMissing
	}
	// Fetch Username
	users, ok := md["user"]
	log := logr.FromContextOrDiscard(ctx)
	if !ok || len(users) == 0 || users[0] == "" {
		log.Info("Username must be set")
		return nil, ErrMetadataMissing
	}
	user := users[0]
	log.Info("USER", "name", user)
	return &user, nil
}
