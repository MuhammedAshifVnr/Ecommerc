package database

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Admin struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"username"`
	Username string `gorm:"unique;not null" json:"useremail"`
	Password string `gorm:"not null" json:"userpassword"`
}

type Product struct {
	gorm.Model
	ProductName  string         `gorm:"not null;unique" json:"prduct"`
	CategoryId   uint           `gorm:"not null"`
	ProductPrice float64        `gorm:"not null" json:"prize"`
	Quantity     int            `gorm:"not null" json:"quantity"`
	Size         int            `gorm:"not null" json:"size"`
	Description  string         `gorm:"not null" json:"description"`
	Status       string         `gorm:"default:Active"`
	ImageUrls    pq.StringArray `gorm:"type:text[]"`
	Category     Category
}

type Category struct {
	gorm.Model
	Name        string `gorm:"unique;not null" json:"catagory"`
	Description string `gorm:"not null" json:"description"`
	Status      string `gorm:"default:Active"`
}

type Review struct {
	gorm.Model
	UserID    uint
	User      User
	ProductID uint
	Product   Product
	Rating    float64
	Comment   string
}
