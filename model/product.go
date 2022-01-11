package model

import (
	"time"
)

type Product struct {
	ID                  int    `gorm:"AUTO_INCREMENT;primary_key"`
	IdHash              string `gorm:"default:NULL"`
	UserId              int    `gorm:"default:NULL"`
	UserIdHash          string `gorm:"default:NULL"`
	ProductName         string `gorm:"default:NULL"`
	Image               string `gorm:"default:NULL"`
	ProductIntroduction string `gorm:"size:4096"`
	ListingFlag         bool   `gorm:"default:1"`
	Price               int    `gorm:"default:NULL"`
	DiscountRate        int    `gorm:"default:0"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           *time.Time
}
