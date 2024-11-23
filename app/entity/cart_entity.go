package entity

import "time"

type Cart struct {
	ID       string     `gorm:"primaryKey" json:"id"`
	Owner    string     `json:"owner"`
	Products []CartItem `gorm:"type:json" json:"products"`
	CreateAt time.Time  `json:"create_at"`
	UpdateAt time.Time  `json:"update_at"`
}
