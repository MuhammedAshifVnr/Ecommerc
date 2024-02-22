package users

import (
	"ecom/database"
	"ecom/helper"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var user database.User
var otp string

type FormOtp struct {
	Userotp string `gorm:"not null" json:"otp"`
}

func Signup(c *gin.Context) {
	user = database.User{}
	var find database.Otp
	c.ShouldBindJSON(&user)
	if err:=helper.DB.Where("email=?",user.Email).First(&find);err==nil{
		c.JSON(200,"Email alredy exist.")
		return
	}
	hashedPassword, error := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if error != nil {
		c.JSON(http.StatusSeeOther, "We can't create hashed password.")
		return
	}

	user.Password = string(hashedPassword)

	otp = helper.GenerateOtp()
	err := helper.DB.Where("email = ?", user.Email).First(&find)
	if err.Error != nil {
		newOtp := database.Otp{
			Secret: otp,
			Expiry: time.Now().Add(10 * time.Second),
			Email:  user.Email,
		}
		helper.DB.Create(&newOtp)
	} else {
		helper.DB.Model(&find).Where("email = ?", user.Email).Updates(database.Otp{
			Secret: otp,
			Expiry: time.Now().Add(60 * time.Second),
		})
	}
	helper.SendOtp(user.Email, otp)

	c.JSON(http.StatusSeeOther, "Otp send the email "+otp)

}

func OtpChecker(c *gin.Context) {
	var formvalue FormOtp
	var find database.Otp
	c.ShouldBindJSON(&formvalue)
	now := time.Now()
	helper.DB.Where("email=? AND expiry > ?", user.Email, now).First(&find)
	fmt.Println(find.Secret)
	if formvalue.Userotp != find.Secret {
		c.JSON(http.StatusSeeOther, "Invalied OTP Please Check You'r Mail.")
		return
	}
	if err := helper.DB.Create(&user); err.Error != nil {
		c.JSON(200, "Email Found Duplicate.")
		return
	}

	c.JSON(200, "Successfully Created User.")

}

func ResendOtp(c *gin.Context) {
	var find database.Otp
	otp = helper.GenerateOtp()
	err := helper.DB.Where("email = ?", user.Email).First(&find)
	if err.Error != nil {
		newOtp := database.Otp{
			Secret: otp,
			Expiry: time.Now().Add(10 * time.Second),
			Email:  user.Email,
		}
		helper.DB.Create(&newOtp)
	} else {
		helper.DB.Model(&find).Where("email = ?", user.Email).Updates(database.Otp{
			Secret: otp,
			Expiry: time.Now().Add(60 * time.Second),
		})
	}
	helper.SendOtp(user.Email, otp)
	c.JSON(200,"send otp to given mail :"+otp)
}
