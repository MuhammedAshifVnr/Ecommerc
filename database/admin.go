package database

import "gorm.io/gorm"

type Admin struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"username"`
	Username string `gorm:"unique;not null" json:"useremail"`
	Password string `gorm:"not null" json:"userpassword"`
}

type Product struct {
	gorm.Model
	ProductName  string `gorm:"not null" json:"prduct"`
	CategoryId   uint   `gorm:"not null"`
	ProductPrize int    `gorm:"not null" json:"prize"`
	Quantity     int    `gorm:"not null" json:"quantity"`
	Size         int    `gorm:"not null" json:"size"`
	Description  string `gorm:"not null" json:"description"`
	Status       string `gorm:"default:Active"`
	Category     Category
}

type Category struct {
	gorm.Model
	Name        string `gorm:"unique;not null" json:"catagory"`
	Description string `gorm:"not null" json:"description"`
	Status      string `gorm:"default:Active"`
}
