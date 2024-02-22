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
}

type Otp struct {
	Secret    string
	Expiry    time.Time
	Email     string
}
