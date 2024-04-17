package users

import (
	"ecom/database"
	"ecom/helper"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// @Summary Signup
// @Description Signup
// @Tags User-Signup
// @Accept  json
// @Produce  json
// @Param user body database.SignupData true "User"
// @Router /user/signup [post]
func Signup(c *gin.Context) {
	var user database.SignupData
	var find database.Otp
	c.ShouldBindJSON(&user)
	if err := helper.DB.Where("email=?", user.Email).First(&find); err == nil {
		c.JSON(200, "Email alredy exist.")
		return
	}
	hashedPassword, error := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if error != nil {
		c.JSON(http.StatusSeeOther, "We can't create hashed password.")
		return
	}

	user.Password = string(hashedPassword)

	otp := helper.GenerateOtp()
	err := helper.DB.Where("email = ?", user.Email).First(&find)
	if err.Error != nil {
		newOtp := database.Otp{
			Secret: otp,
			Expiry: time.Now().Add(120 * time.Second),
			Email:  user.Email,
		}
		helper.DB.Create(&newOtp)
	} else {
		helper.DB.Model(&find).Where("email = ?", user.Email).Updates(database.Otp{
			Secret: otp,
			Expiry: time.Now().Add(60 * time.Second),
		})
	}
	data := map[string]interface{}{
		"name":     user.Name,
		"email":    user.Email,
		"password": hashedPassword,
	}
	if err:=helper.SendOtp(user.Email, otp);err!=nil{
		c.JSON(400,gin.H{"code":400,"status":"Error","message":""})
	}
	// session
	session := sessions.Default(c)
	session.Set(user.Email, data)
	session.Save()
	c.SetCookie("sessionID", user.Email, 3600, "/", "ashif.online", false, false)

	c.JSON(http.StatusSeeOther, "Otp send the email "+otp)

}

// @Summary OtpChecker
// @Description OtpChecker
// @Tags User-Signup
// @Accept  multipart/form-data
// @Produce  json
// @Param OTP formData string true "OTP"
// @Router /user/otp [post]
func OtpChecker(c *gin.Context) {
	var find database.Otp
	otp := c.Request.FormValue("OTP")
	now := time.Now()
	helper.DB.Where("secret=? AND expiry > ?", otp, now).First(&find)
	fmt.Println(find.Secret)
	if otp != find.Secret {
		c.JSON(http.StatusSeeOther, "Invalied OTP Please Check You'r Mail.")
		return
	}
	session := sessions.Default(c)
	cookie, _ := c.Cookie("sessionID")
	userDatas := session.Get(cookie)
	userMap, _ := userDatas.(map[string]interface{})
	fmt.Println(userMap)
	if err := helper.DB.Create(&database.User{
		Name:     userMap["name"].(string),
		Email:    userMap["email"].(string),
		Password: string(userMap["password"].([]uint8)),
	}); err.Error != nil {
		c.JSON(200, "Email Found Duplicate.")
		return
	} else {
		var user database.User
		if err := helper.DB.Where("email=?", find.Email).First(&user); err.Error != nil {
			fmt.Println("email fount :", find.Email)
			fmt.Println("==========", err.Error)
		}
		wallet := database.Wallet{
			UserID: user.ID,
			Amount: 0.0,
		}
		helper.DB.Create(&wallet)
		c.JSON(200, "Successfully Created User.")
	}

}

// @Summary Resend OTP
// @Description Resend OTP
// @Tags User-Signup
// @Accept  json
// @Produce  json
// @Router /user/reotp [post]
func ResendOtp(c *gin.Context) {
	var find database.Otp
	var user database.User
	otp := helper.GenerateOtp()
	email,_:=c.Cookie("sessionID")
	fmt.Println("======",email)
	err := helper.DB.Where("email = ?", email).First(&find)
	if err.Error != nil {
		newOtp := database.Otp{
			Secret: otp,
			Expiry: time.Now().Add(10 * time.Second),
			Email:  user.Email,
		}
		helper.DB.Create(&newOtp)
	} else {
		helper.DB.Model(&find).Where("email = ?", email).Updates(database.Otp{
			Secret: otp,
			Expiry: time.Now().Add(60 * time.Second),
		})
	}
	helper.SendOtp(user.Email, otp)
	c.JSON(200, "send otp to given mail :"+otp)
}
