package entity

import "time"

type Cart struct {
	ID       string     `gorm:"primaryKey" json:"id"`
	Products string     `json:"products"`
	CreateAt *time.Time `json:"create_at"`
	UpdateAt *time.Time `json:"update_at"`
}
