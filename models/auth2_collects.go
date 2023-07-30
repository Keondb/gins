package models

import "time"

type User2Collects struct {
	UserID       uint         `gorm:"primaryKey"`
	UserModel    UserModel    `gorm:"foreignKey:UserID" json:"-"`
	ArticleModel ArticleModel `gorm:"foreignKey:ArticleID" json:"-"`
	ArticleID    uint         `gorm:"primaryKey"`
	CreatedAt    time.Time
}
