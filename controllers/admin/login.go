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
		token, err := middleware.GenerateToken("admin", AdminTable.Username, AdminTable.ID, AdminTable.Name)
		if err != nil {
			fmt.Println("TOken cant generate.")
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
		"Message":           "Welcome " + admin,
		"TotalSaleAmount ": total,
		"ConmformCount ":   conform,
		"CancelledOrders ": cancel,
	})

	// for _, order := range orders {
	// 	c.JSON(303, gin.H{
	// 		"ProductName":   order.Product.ProductName,
	// 		"OrderID":       order.OrderID,
	// 		"Amount":        order.Amount,
	// 		"ID":            order.ID,
	// 		"PaymentMethod": order.Order.PaymentMethod,
	// 		"UserName":      order.Order.User.Email,
	// 		"Quantity":      order.Quantity,
	// 		"Status":        order.Status,
	// 	})
	// }

}

func Logout(c *gin.Context) {

	c.SetCookie("admin", "", -1, "/", "localhost", false, true)
	c.JSON(200, "Successfully logout.")
}
