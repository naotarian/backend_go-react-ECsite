package model

import "time"

type User struct {
	ID        int `gorm:"AUTO_INCREMENT;primary_key"`
	IdHash    *string
	Name      *string
	Icon      *string
	Gender    *string
	Email     *string `gorm:"type:varchar(100);unique_index"`
	Comment   *string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
