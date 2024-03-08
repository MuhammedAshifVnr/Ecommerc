package users

import (
	"ecom/database"
	"ecom/helper"
	"ecom/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var roleUser = "user"
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
	if Find.Status=="Block"{
		c.JSON(200, "User Blocked")
		return
	}

	middleware.SessionCreate(Find.Email, roleUser, c)

	c.JSON(200, "Sccessfully Logedin.")

}

func Logout(c *gin.Context) {

	session := sessions.Default(c)
	check := session.Get("user")
	if check == nil {
		c.JSON(200, "Not logged in")
	} else {
		Find = database.User{}
		session.Delete("user")
		session.Save()
		c.JSON(200, "Successfully logout.")
	}
}
