package users

import (
	"ecom/database"
	"ecom/helper"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)
var Find database.User
func Login(c *gin.Context) {
	user = database.User{}
	
	c.ShouldBindJSON(&user)
	helper.DB.First(&Find, "email=?", user.Email)

	err := bcrypt.CompareHashAndPassword([]byte(Find.Password), []byte(user.Password))
	if err != nil {
		c.JSON(200, "Invalid Username or Password.")
		return
	}
	c.JSON(200, "Sccessfully Logedin.")

}

func Logout (c *gin.Context){
	Find=database.User{}
	c.JSON(200,"Successfully logout.")
}