package users

import (
	"ecom/database"
	"ecom/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CheckOut(c *gin.Context) {
	var cartItems []database.Cart
	helper.DB.Preload("Product").Where("user_id=?", Find.ID).Find(&cartItems)

	paymentMethod := c.Request.FormValue("payment")
	Address, _ := strconv.ParseUint(c.Request.FormValue("address"), 10, 64)

	if paymentMethod == "" || Address == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Payment Method and Address are required",
		})
		return
	}

	couponCode := c.Request.FormValue("coupon")
	var coupon database.Coupon
	if couponCode != "" {
		if err := helper.DB.Where("code=?", couponCode).First(&coupon); err.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid coupon code",
			})
			return
		}
	}
	var totalAmount float64
	for _, cartItem := range cartItems {
		Amount := (cartItem.Product.ProductPrice * float64(cartItem.Quantity))

		if cartItem.Quantity > uint(cartItem.Product.Quantity) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Insufficent stock for product " + cartItem.Product.ProductName,
			})
			return
		}

		cartItem.Product.Quantity -= int(cartItem.Quantity)
		if err := helper.DB.Save(&cartItem.Product).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Faild to Update Product Stock",
			})
			return
		}

		order := database.Order{
			UserID:        Find.ID,
			PaymentMethod: paymentMethod,
			AddressID:     uint(Address),
			ProductID:     cartItem.Product.ID,
			Quantity:      cartItem.Quantity,
		}
		if couponCode != "" {
			Amount -= coupon.Amount
			order.CouponID = coupon.ID
		} else {
			order.CouponID = 4
		}
		order.Amount = Amount
		helper.DB.Create(&order)
		totalAmount += Amount
	}

	if err := helper.DB.Where("user_id =?", Find.ID).Delete(&database.Cart{}); err.Error != nil {
		c.JSON(http.StatusBadRequest, "faild to delete datas in cart.")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order Placed Successfully.",
		"Amount":  totalAmount,
	})

}

func Order(c *gin.Context) {
	var orders []database.Order
	helper.DB.Where("user_id=?", Find.ID).Find(&orders)
	for _, order := range orders {
		c.JSON(200, gin.H{
			"ID":      order.ID,
			"Product": order.Product,
			"Amount":  order.Amount,
			"Status":  order.Status,
		})
	}
}

func OrderDetils(c *gin.Context) {
	var order database.Order
	id := c.Param("ID")
	helper.DB.Preload("Coupon").Where("id=?", id).First(&order)
	c.JSON(200, gin.H{
		"Product":         order.Product,
		"Amount":          order.Amount,
		"Coupon":          order.Coupon.Code,
		"Status":          order.Status,
		"Payment Method":  order.PaymentMethod,
		"Order Confirmed": order.Model.CreatedAt,
		"Status Updated":  order.Model.UpdatedAt,
	})
}

func CancelOrder(c *gin.Context) {
	id := c.Param("ID")
	var order database.Order
	helper.DB.Where("id=?", id).First(&order)
	order.Status = "cancelled"
	order.Reason = c.Request.FormValue("reason")
	helper.DB.Save(&order)
	c.JSON(200, "Order Cancelled.")
}
