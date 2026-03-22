type Container struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Location    string         `json:"location" gorm:"not null"`
	City        string         `json:"city" gorm:"default:'Paris'"`
	Capacity    int            `json:"capacity"`
	Status      string         `json:"status" gorm:"default:'active'"`
	LastEmptied time.Time      `json:"last_emptied"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdateAt    time.Time      `json:"updated_at"`
	DeleteAt    gorm.DeletedAt `json:"-" gorm:"index"`
}
