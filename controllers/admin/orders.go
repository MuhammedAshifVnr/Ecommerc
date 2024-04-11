package admin

import (
	"ecom/database"
	"ecom/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary      Orders Listing
// @Description  All Orders  are listed here
// @Tags Admin-Order
// @Produce      json
// @Router       /admin/order [get]
func Orders(c *gin.Context) {
	var orders []database.OrderItems
	helper.DB.Preload("Product").Preload("Order.User").Preload("Order.Coupon").Find(&orders)
	var response []gin.H
	for _, order := range orders {
		response = append(response, gin.H{
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
	c.JSON(http.StatusOK, gin.H{
		"Status": http.StatusOK,
		"Data":   response,
	})
}

// UpdateOrder updates the status of an order.
// @Summary Update Order Status
// @Description Update the status of an order by ID.
// @ID update-order
// @Tags Admin-Order
// @Accept multipart/form-data
// @Produce json
// @Param ID path int true "Order ID"
// @Param status formData string true "New status"
// @Router       /admin/order/update/{ID} [patch]
func UpdateOrder(c *gin.Context) {
	id := c.Param("ID")
	var order database.OrderItems
	helper.DB.Where("id=?", id).First(&order)
	if order.Status == "Cancelled" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Order alredy Cancelled",
		})
		return
	}
	order.Status = c.Request.FormValue("status")
	helper.DB.Save(&order)
	c.JSON(200, gin.H{
		"Message": "Order Status Updated",
	})

}
