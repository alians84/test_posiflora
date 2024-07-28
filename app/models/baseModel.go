package models

import "time"

type BaseModel struct {
	ID         uint64    `gorm:"primary_key" json:"id"`
	DateCreate time.Time `json:"date_create"`
	DateUpdate time.Time `json:"date_update"`
}
