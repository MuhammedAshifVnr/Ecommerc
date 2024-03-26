package users

import (
	"ecom/database"
	"ecom/helper"
	"ecom/middleware"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	user = database.User{}
	var Find database.User
	c.ShouldBindJSON(&user)
	helper.DB.First(&Find, "email=?", user.Email)

	err := bcrypt.CompareHashAndPassword([]byte(Find.Password), []byte(user.Password))
	if err != nil {
		c.JSON(200, "Invalid Username or Password.")
		return
	}
	if Find.Status == "Block" {
		c.JSON(200, "User Blocked")
		return
	}

	token, err := middleware.GenerateToken(Find.Role, Find.Email, Find.ID,Find.Name)
	if err != nil {
		fmt.Println("TOken cant generate.")
	}
	c.SetCookie("accessToken", token, int((time.Hour * 24).Seconds()), "/", "localhost", false, true)
	fmt.Println(token)
	c.JSON(200, gin.H{
		"messe":"successfully Login.",
		"token":token,
	})

}

func Logout(c *gin.Context) {
	c.SetCookie("accessToken", "", -1, "/", "localhost", false, true)
	c.JSON(200, "Successfully logout.")

}
