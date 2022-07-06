package service

import (
	"time"
)

type Status int

const (
	StatusWait    Status = 1
	StatusRead           = 2
	StatusExpired        = 3
	StatusCleared        = 4
)

type Item struct {
	Title string `json:"title"`
	Group string `json:"group"`
}

type NewItemRequest struct {
	Item
	ExpireValue   string `json:"exp"`
	ExpireInHours bool   `json:"exp_hours"`
	Data          string `json:"data"`
}

type ItemMeta struct {
	Item
	Owner    string    `json:"owner"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"` // >now() means Expire (and Status=1)
	Status   Status    `json:"status"`
}

type ItemInfo struct {
	Id   string    `json:"id"`
	Meta *ItemMeta `json:"meta"`
}

type Stat struct {
	Total   int `json:"total"`
	Wait    int `json:"wait"`
	Read    int `json:"read"`
	Expired int `json:"expired"`
}

type StatResponse struct {
	My    Stat `json:"my"`
	Other Stat `json:"other"`
}
