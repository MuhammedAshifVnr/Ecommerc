package users

import (
	"ecom/database"
	"ecom/helper"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	user = database.User{}
	var find database.User
	c.ShouldBindJSON(&user)
	helper.DB.First(&find, "email=?", user.Email)

	err := bcrypt.CompareHashAndPassword([]byte(find.Password), []byte(user.Password))
	if err != nil {
		c.JSON(200, "Invalid Username or Password.")
		return
	}
	c.JSON(200, "Sccessfully Logedin.")

}
