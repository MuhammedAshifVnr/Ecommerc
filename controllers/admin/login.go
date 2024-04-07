package admin

import (
	"ecom/database"
	"ecom/helper"
	"ecom/middleware"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// ..................................login post......................
// AdminLogin godoc
// @Summary      Login an admin user
// @Description  Login an admin user with username and password
// @Tags Admin
// @ID           admin-login
// @Accept       json
// @Produce      json
// @Param        admin body database.Logging true "Admin login details"
// @Router       /admin/login [post]
func AdminLogin(c *gin.Context) {
	var find database.Logging
	var AdminTable database.Admin
	if err := c.ShouldBindJSON(&find); err != nil {
		c.JSON(http.StatusSeeOther, "Something went wrong.")
	}

	helper.DB.First(&AdminTable, "Username=?", find.Username)

	if AdminTable.Password != find.Password {
		c.JSON(401, "Invalid Username or Password")
		fmt.Println(AdminTable, find)
	} else {
		token, err := middleware.GenerateToken("admin", AdminTable.Username, AdminTable.ID, AdminTable.Name)
		if err != nil {
			c.JSON(403, gin.H{"error": "failed to create token"})
			return
		}
		c.SetCookie("admin", token, int((time.Hour * 24).Seconds()), "/", "localhost", false, true)
		fmt.Println(token)
		c.JSON(200, gin.H{
			"massege": "Successfuly logined",
			"token":   token,
		})

	}
}

//..............after login show this page list of users.........................
// AdminLogin godoc
// @Summary      Admin Home Page
// @Description  after login show this page 
// @Tags Admin
// @Produce      json
// @Router       /admin/home [get]
func HomePage(c *gin.Context) {
	admin := c.GetString("username")
	var orders []database.OrderItems
	helper.DB.Preload("Product").Preload("Order.User").Find(&orders)
	var total float64
	var cancel, conform int
	for _, order := range orders {
		if order.Status == "Cancelled" {
			cancel++
		} else {
			conform++
		}
		total += order.Amount
	}
	c.JSON(http.StatusSeeOther, gin.H{
		"Message":          "Welcome " + admin,
		"TotalSaleAmount ": total,
		"ConmformCount ":   conform,
		"CancelledOrders ": cancel,
	})

}

// AdminLogin godoc
// @Summary      Logout an admin user
// @Description  Logout an admin clearing the cookie
// @Tags Admin
// @Produce      json
// @Router       /admin/logout [get]
func Logout(c *gin.Context) {

	c.SetCookie("admin", "", -1, "/", "localhost", false, true)
	c.JSON(200, "Successfully logout.")
}
