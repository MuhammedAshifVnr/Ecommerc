package admin

import (
	"ecom/database"
	"ecom/helper"
	"ecom/jwt"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// ..................................login post......................

func AdminLogin(c *gin.Context) {
	var find database.Admin
	var AdminTable database.Admin
	if err := c.ShouldBindJSON(&find); err != nil {
		c.JSON(http.StatusSeeOther, "Something went wrong.")
	}

	helper.DB.First(&AdminTable, "Username=?", find.Username)

	if AdminTable.Password != find.Password {
		c.JSON(http.StatusSeeOther, "Invalid Username or Password")
		fmt.Println(AdminTable, find)
	} else {
		token, err := jwt.GenerateToken("admin", AdminTable.Username, AdminTable.ID, AdminTable.Name)
		if err != nil {
			fmt.Println("TOken cant generate.")
		}
		fmt.Println(token)
		c.JSON(200, "Successfuly logined")

	}
}

//..............after login show this page list of users.........................

func HomePage(c *gin.Context) {
	admin := c.GetString("username")
	c.JSON(http.StatusSeeOther, gin.H{
		"Message": "Welcome " + admin,
	})

}

func Logout(c *gin.Context) {

	session := sessions.Default(c)
	check := session.Get("admin")
	if check == nil {
		c.JSON(200, "Not logged in")
	} else {
		session.Delete("admin")
		session.Save()
		c.JSON(200, "Successfully logout.")
	}
}
