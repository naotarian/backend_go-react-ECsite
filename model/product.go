package model

import (
	"time"
)

type Product struct {
	ID                  int `gorm:"AUTO_INCREMENT;primary_key"`
	IdHash              *string
	UserId              *int
	UserIdHash          *string
	ProductName         *string
	Image               *string
	ProductIntroduction *string `gorm:"size:4096"`
	ListingFlag         bool    `gorm:"default:1"`
	Price               int
	DiscountRate        *int `gorm:"default:0"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           *time.Time
}
