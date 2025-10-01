package entity

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement"`
	AuthorID    uint64 `gorm:"not null"`
	Author      User   `gorm:"foreignKey:AuthorID"`
	Title       string `gorm:"size:255;not null"`
	Content     string
	IsPublished bool `gorm:"default:false"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
