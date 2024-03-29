package admin

import (
	"ecom/database"
	"ecom/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Orders(c *gin.Context) {
	var orders []database.OrderItems
	helper.DB.Preload("Product").Preload("Order.User").Preload("Order.Coupon").Find(&orders)
	for _, order := range orders {
		c.JSON(200, gin.H{
			"ID":          order.ID,
			"OrderID":     order.OrderID,
			"ProductName": order.Product.ProductName,
			"Coupon":      order.Order.Coupon.Code,
			"Quantity":    order.Quantity,
			"Amount":      order.Amount,
			"User":        order.Order.User.Email,
			"Status":      order.Status,
		})
	}
}

func UpdateOrder(c *gin.Context) {
	id := c.Param("ID")
	var order database.OrderItems
	helper.DB.Where("id=?", id).First(&order)
	if order.Status == "Cancelled" {
		c.JSON(http.StatusBadRequest, gin.H{
			"massege": "Order alredy Cancelled",
		})
		return
	}
	order.Status = c.Request.FormValue("status")
	helper.DB.Save(&order)
	c.JSON(200, gin.H{
		"massege": "Order Status Updated",
	})

}
