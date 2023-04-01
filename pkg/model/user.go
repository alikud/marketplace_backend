package model

import "time"

type User struct {
	Id              int64     `json:"id,omitempty"`
	Email           string    `json:"email,omitempty"`
	Password        string    `json:"password,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	ShippingAddress string    `json:"shipping_address,omitempty"`
}
