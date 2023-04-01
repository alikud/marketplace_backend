package model

import "time"

type StoreProduct struct {
	ProductId      int64     `json:"product_id,omitempty"`
	StoreId        int64     `json:"store_id,omitempty"`
	Price          float64   `json:"price,omitempty"`
	AvailableCount int       `json:"available_count,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
}
