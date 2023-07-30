// DONT EDIT: Auto generated by ifacemaker

// Package storage hold Iface which all subpackages implements.
package storage

import (
	gen "github.com/LeKovr/showonce/zgen/go/proto"
	"github.com/oklog/ulid/v2"
)

// Iface makes users independent from storage implementation.
type Iface interface {
	// SetMeta prepares and saves item metadata and secret.
	SetMeta(owner string, req *gen.NewItemRequest) (*ulid.ULID, error)
	// GetMeta returns item metadata.
	GetMeta(id string) (*gen.ItemMeta, error)
	// GetData returns item data (secret).
	GetData(id string) (*gen.ItemData, error)
	// Items returns items, created by current user.
	Items(owner string) (*gen.ItemList, error)
	// Stats returns global and user's item counters.
	Stats(owner string) (*gen.StatsResponse, error)
}