// Package storage holds common data for dependent packages.
package storage

import "errors"

var (
	// ErrNoUniqueWithinLimit means we cant find unique ID within given time period.
	ErrNoUniqueWithinLimit = errors.New("cannot create unique id")
	// ErrNotFound means we have no data gor given ID.
	ErrNotFound = errors.New("item not found")
	// ErrDataCorrupted means we cannot fetch data from cache.
	ErrDataCorrupted = errors.New("item data was corrupted")
)
