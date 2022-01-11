package model

import "time"

type User struct {
	ID        int    `gorm:"AUTO_INCREMENT;primary_key"`
	IdHash    string `gorm:"default:NULL"`
	Name      string `gorm:"default:NULL"`
	Icon      string `gorm:"default:NULL"`
	Gender    string `gorm:"default:NULL"`
	Email     string `gorm:"type:varchar(100);unique_index;default:NULL"`
	Comment   string `gorm:"size:255;default:NULL"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
