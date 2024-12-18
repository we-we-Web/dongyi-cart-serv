package domain

import "time"

type Cart struct {
	ID       string     `json:"id"`
	Products []CartItem `json:"products"`
	CreateAt time.Time  `json:"create_at"`
	UpdateAt time.Time  `json:"update_at"`
}

type CartItem struct {
	Product string         `json:"product"`
	Spec    map[string]int `json:"spec"`
}

func NewCart(id string, t time.Time) *Cart {
	return &Cart{
		ID:       id,
		Products: []CartItem{},
		CreateAt: t,
		UpdateAt: t,
	}
}

func NewCartItem(productID string, size string, quantity int) *CartItem {
	return &CartItem{
		Product: productID,
		Spec: map[string]int{
			size: quantity,
		},
	}
}
