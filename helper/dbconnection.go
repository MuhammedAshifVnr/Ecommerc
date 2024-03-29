package helper

import (
	"ecom/database"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func EnvLoader() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("-------------Faild to load env file------------ ")
	}
}

func DbConnect() {

	dsn := os.Getenv("DSN")

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		fmt.Println(".....................Connection faild..................")
	}

	DB = db

	DB.AutoMigrate(&database.Wallet{}, &database.OrderItems{}, &database.Order{}, &database.User{}, &database.Admin{}, &database.Category{}, &database.Otp{})
	DB.AutoMigrate(&database.Offers{}, &database.Transactions{}, &database.Whislist{}, &database.Coupon{}, &database.Cart{}, &database.Address{}, &database.Product{}, &database.Review{})
	fmt.Println("SUCCESSFULLY connected to DATABASE")

}
