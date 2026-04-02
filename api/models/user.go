package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             uint           `json: "id" gorm: "primarykey"`
	Nom            string         `json: "nom" gorm: not null`
	Email          string         `json: "email" gorm:"unique; not null"`
	Password       string         `json: "password,omitempy" gorm: "not null"`
	Role           string         `json: "role" gorm: "default: 'individual'"`
	AvatarURL      string         `json: "avaltar_url"`
	UpcyclingScore int            `json: "upcyclingScore"`
	CreateAt       time.Time      `json: "created_at"`
	UpdatedAt      time.Time      `json:"updated_at`
	DeletedAt      gorm.DeletedAt `json:"-" gorm:"index"`
}
