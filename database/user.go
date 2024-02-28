package database

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null" json:"username"`
	Email    string `gorm:"unique;not null" json:"useremail"`
	Password string `gorm:"not null" json:"userpassword"`
	Status   string `gorm:"default:Active"`
	Mobile   int    `json:"mobile"`
	Gender   string `json:"gender"`
	Address  Address
}

type Otp struct {
	Secret string
	Expiry time.Time
	Email  string
}

type Address struct {
	ID      uint `gorm:"primarykey"`
	UserId  uint
	Type    string `jsonz:"type"`
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode uint   `json:"zip"`
}

type Cart struct {
	ID        uint
	UserID    uint
	ProductID uint
	Product   Product
	Quantity  uint
	SubTotal  uint
}
