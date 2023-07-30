package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primarykey" json:"id"` // 主键Id
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}
