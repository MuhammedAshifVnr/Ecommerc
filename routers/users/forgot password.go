package users

import (
	"ecom/database"
	"ecom/helper"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var forgot database.User

func EmailChecking(c *gin.Context) {
	forgot = database.User{}
	var check database.Otp

	if err := helper.DB.Where("email=?", c.Request.FormValue("email")).First(&forgot); err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email not found.",
		})
		return
	}
	otp := helper.GenerateOtp()
	helper.DB.Model(&check).Where("email = ?", forgot.Email).Updates(database.Otp{
		Secret: otp,
		Expiry: time.Now().Add(60 * time.Second),
	})
	helper.SendOtp(forgot.Email, otp)
	c.JSON(200, "Otp sended to email."+otp)
}

func OtpValidation(c *gin.Context) {

	var find database.Otp
	now := time.Now()
	helper.DB.Where("email=? AND expiry>?", forgot.Email, now).First(&find)
	if c.Request.FormValue("otp") != find.Secret {
		c.JSON(http.StatusBadRequest, "Invalied OTP Please Check You'r Mail.")
		return
	}
	c.JSON(200, "Enter New Password.")
}

func UpdatePassword(c *gin.Context) {
	if c.Request.FormValue("one") != c.Request.FormValue("two") {
		c.JSON(http.StatusBadRequest, "Password not maching.")
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(c.Request.FormValue("one")), 8)
	if err != nil {
		c.JSON(http.StatusSeeOther, "We can't create hashed password.")
		return
	}
	helper.DB.Model(&forgot).Where("email=?", forgot.Email).Update("Password", string(hashedPassword))
	c.JSON(200, "Password Updated Successfully.")
}
