package model

import "time"

type Product struct {
	Id          int64     `json:"id,omitempty"`
	CategoryId  string    `json:"category_id,omitempty"`
	Name        string    `json:"name,omitempty"`
	BrandId     string    `json:"brand_id,omitempty"`
	Description string    `json:"description,omitempty"`
	PictureUrl  string    `json:"picture_url,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}
