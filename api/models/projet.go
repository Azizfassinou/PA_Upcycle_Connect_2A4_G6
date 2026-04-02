package models

import (
	"time"

	"gorm.io/gorm"
)

type Projet struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"not null"`
	Description string         `json:"description"`
	Category    string         `json:"category"`
	Difficulty  string         `json:"difficulty"`
	Author_ID   uint           `json:"author_id"`
	Author      User           `json:"author" gorm:"foreignKey:Author_ID"`
	Steps       string         `json:"steps" gorm:"type:text"`
	ImageURL    string         `json:"image_url"`
	Status      string         `json:"status" gorm:"default:'completed'"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeleteAt    gorm.DeletedAt `json:"-" gorm:"index"`
}
