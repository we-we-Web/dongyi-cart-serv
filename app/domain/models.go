package domain

import "time"

type Cart struct {
	ID       string     `json:"id"`
	Products []CartItem `json:"products"`
	CreateAt *time.Time `json:"create_at"`
	UpdateAt *time.Time `json:"update_at"`
}

type CartItem struct {
	Product  string `json:"product"`
	Quantity int    `json:"quantity"`
}

func NewCart(id string, t *time.Time) *Cart {
	return &Cart{
		ID:       id,
		Products: []CartItem{},
		CreateAt: t,
		UpdateAt: t,
	}
}

func NewCartItem(productID string, quantity int) *CartItem {
	return &CartItem{
		Product:  productID,
		Quantity: quantity,
	}
}
