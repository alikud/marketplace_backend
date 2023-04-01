package model

import "time"

type OrderProduct struct {
	ProductId int64     `json:"product_id,omitempty"`
	StoreId   int64     `json:"store_id,omitempty"`
	OrderId   int64     `json:"order_id,omitempty"`
	Quantity  int64     `json:"quantity,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}
