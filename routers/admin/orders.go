package admin

import (
	"ecom/database"
	"ecom/helper"

	"github.com/gin-gonic/gin"
)

func Orders(c *gin.Context) {
	var orders []database.Order

	helper.DB.Preload("User").Preload("Product").Preload("Coupon").Find(&orders)

	// for _, v := range orders {
	// 	c.JSON(200, gin.H{
	// 		"Id":       v.ID,
	// 		"Username": v.User.Email,
	// 		"Product": v.Product.ProductName,
	// 		"Quantity":v.Quantity,
	// 		"Price":v.Amount,
	// 		"Coupon":v.Coupon.Code,
	// 		"Status":v.Status,
	// 		"Reason":v.Reason,
	// 	})
	// }
}

func UpdateOrder(c *gin.Context)  {
	id := c.Param("ID")
	var order database.Order
	helper.DB.Preload("Product").Where("id=?", id).First(&order)
	// order.Status = c.Request.FormValue("status")
	// if order.Status ==""{
	// 	c.JSON(400,"Status Field Required")
	// }
	helper.DB.Save(&order)
	c.JSON(200,"Status Updated.")
}