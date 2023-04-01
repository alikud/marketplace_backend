package model

import "time"

type Store struct {
	Id          int64     `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	OwnerId     int64     `json:"owner_id,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}
