package models

import "time"

type Paste struct {
	Url            string        `json:"url"`
	Title          string        `json:"title"`
	Data           string        `json:"data"`
	Status         string        `json:"status"`
	ExpirationTime time.Duration `json:"expiration_time"`
	CreateTime     time.Time     `json:"-"`
	UserId         int64         `json:"-"`
}
