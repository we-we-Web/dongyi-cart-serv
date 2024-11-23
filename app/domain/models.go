package domain

import "time"

type Cart struct {
	ID       string     `json:"id"`
	Owner    string     `json:"owner"`
	Products []CartItem `json:"products"`
	CreateAt *time.Time `json:"create_at"`
	UpdateAt *time.Time `json:"update_at"`
}

type CartItem struct {
	Product  string `json:"product"`
	Quantity int    `json:"quantity"`
}

func NewCart(id, owner string, t *time.Time) *Cart {
	return &Cart{
		ID:       id,
		Owner:    owner,
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
