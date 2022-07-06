package models

import "time"

type MyModel struct {
	Id        uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	CreatedAt MyTime `gorm:"column:created_at" json:"created_at"`
	UpdatedAt MyTime `gorm:"column:updated_at" json:"updated_at"`
}

type MyTime struct {
	time.Time
}
