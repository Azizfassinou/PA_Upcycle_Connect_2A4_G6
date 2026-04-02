package models

import (
	"time"

	"gorm.io/gorm"
)

type Announcement struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"not null"`
	Description string         `json:"description"`
	Category    string         `json:"category"`
	Type        string         `json:"type"`
	AuthorID    uint           `json:"author_id"`
	Author      User           `json:"author" gorm:"foreignKey:AuthorID"`
	Status      string         `json:"status" gorm:"default:'available'"`
	ImageURL    string         `json:"image_url"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeleteAt    gorm.DeletedAt `json:"-" gorm:"index"`
}
