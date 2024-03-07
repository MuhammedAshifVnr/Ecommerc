package admin

import (
	"ecom/database"
	"ecom/helper"

	"github.com/gin-gonic/gin"
)

func Orders(c *gin.Context) {
	var orders []database.Order

	helper.DB.Preload("User").Preload("Product").Preload("Coupon").Find(&orders)

	for _, v := range orders {
		c.JSON(200, gin.H{
			"Id":       v.ID,
			"Username": v.User.Email,
			"Product": v.Product.ProductName,
			"Quantity":v.Quantity,
			"Price":v.Amount,
			"Coupon":v.Coupon.Code,
		})
	}
}
