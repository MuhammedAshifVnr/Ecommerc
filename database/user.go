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
	Role     string `gorm:"default:user"`
}

type Otp struct {
	Secret string
	Expiry time.Time
	Email  string
}

type Address struct {
	ID      uint `gorm:"primarykey"`
	UserId  uint
	User    User
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
}

type Coupon struct {
	gorm.Model
	Code   string `gorm:"unique"`
	Amount float64
}

type Order struct {
	ID            uint `gorm:"primarykey"`
	UserID        uint `gorm:"not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	User          User
	PaymentMethod string `gorm:"not null"`
	AddressID     uint   `gorm:"not null"`
	Address       Address
	CouponID      uint `gorm:"default:4"`
	Coupon        Coupon
	Amount        float64
}

type OrderItems struct {
	gorm.Model
	ProductID uint
	Product   Product
	Quantity  uint
	Amount    float64
	OrderID   uint
	Order     Order
	Status    string `gorm:"not null;default:'pending'"`
	Reason    string
}

type Whislist struct{
	gorm.Model
	UserID uint
	User User
	ProductID uint
	Product Product
}