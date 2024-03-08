package admin

import (
	"ecom/database"
	"ecom/helper"
	"ecom/middleware"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var AdminTable database.Admin

// ..................................login post......................

func AdminLogin(c *gin.Context) {
	var find database.Admin
	if err := c.ShouldBindJSON(&find); err != nil {
		c.JSON(http.StatusSeeOther, "Something went wrong.")
	}

	helper.DB.First(&AdminTable, "Username=?", find.Username)

	if AdminTable.Password != find.Password {
		c.JSON(http.StatusSeeOther, "Invalid Username or Password")
		fmt.Println(AdminTable, find)
	} else {
		middleware.SessionCreate(find.Username,"admin",c)
		c.JSON(200, "Successfuly logined")

	}
}

//..............after login show this page list of users.........................

func HomePage(c *gin.Context) {

	c.JSON(http.StatusSeeOther, "Welcome "+AdminTable.Name)

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
