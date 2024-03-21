package users

import (
	"crypto/rand"
	"ecom/database"
	"ecom/helper"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Testcheckout(c *gin.Context) {
	var cartitems []database.Cart
	userId := c.GetUint("userID")
	helper.DB.Preload("Product").Where("user_id=?", userId).Find(&cartitems)
	paymentMethod := c.Request.FormValue("payment")
	Address, _ := strconv.ParseUint(c.Request.FormValue("address"), 10, 64)
	couponCode := c.Request.FormValue("coupon")

	if paymentMethod == "" || Address == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Payment Method and Address are required",
		})
		return
	}

	var coupon database.Coupon
	if couponCode != "" {
		if err := helper.DB.Where("code=?", couponCode).First(&coupon); err.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid coupon code",
			})
			return
		}
	}
	tx := helper.DB.Begin()

	orderID, _ := strconv.ParseUint(generateRandomNumber(), 10, 64)
	order := database.Order{
		UserID:        userId,
		AddressID:     uint(Address),
		PaymentMethod: paymentMethod,
		CouponID:      coupon.ID,
		ID:            uint(orderID),
	}
	if err := tx.Create(&order); err.Error != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}
	var totalAmount float64
	for _, item := range cartitems {
		amount := item.Product.ProductPrice * float64(item.Quantity)
		if item.Quantity > uint(item.Product.Quantity) {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Insufficent stock for product " + item.Product.ProductName,
			})
			return
		}

		item.Product.Quantity -= int(item.Quantity)
		if err := helper.DB.Save(&item.Product).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Faild to Update Product Stock",
			})
			return
		}

		orderItem := database.OrderItems{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Amount:    amount,
			OrderID:   uint(orderID),
		}
		if err := tx.Create(&orderItem); err.Error != nil {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Items cant cerate.",
			})
			return
		}
		totalAmount += amount
	}
	if totalAmount <= float64(coupon.Limit) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You can't use this coupon."})
		return
	}
	totalAmount -= coupon.Amount

	tx.Model(&database.Order{}).Where("id=?", orderID).Update("amount", totalAmount)
	tx.Commit()

	if order.PaymentMethod == "ONLINE" {
		paymentOrder, err := razorpayClient.Order.Create(map[string]interface{}{
			"amount":   totalAmount * 100,
			"currency": "INR",
			"receipt":  strconv.Itoa(int(orderID)),
		}, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		helper.DB.Model(&database.Order{}).Where("id=?", orderID).Update("payment_id", paymentOrder["id"])
		c.JSON(http.StatusOK, gin.H{
			"message":   "Order Placed Successfully.",
			"Amount":    totalAmount,
			"PaymentID": paymentOrder["id"],
		})
	} else if order.PaymentMethod == "COD" {
		c.JSON(http.StatusOK, gin.H{
			"message": "Order Placed Successfully. Product Get Soon",
			"Amount":  totalAmount,
		})
	}

	//=========must implement after checkout cartitems delete======

}

func Order(c *gin.Context) {
	var orders []database.OrderItems
	userId := c.GetUint("userID")
	helper.DB.Preload("Order", "user_id=?", userId).Preload("Product").Find(&orders)
	for _, order := range orders {
		c.JSON(200, gin.H{
			"ID":        order.ID,
			"ProductID": order.Product.ID,
			"Product":   order.Product.ProductName,
			"OrderID":   order.Order.ID,
		})
	}
}

func OrderDetils(c *gin.Context) {
	var orders database.OrderItems
	id := c.Param("ID")
	helper.DB.Preload("Order").Preload("Product").Preload("Order.Coupon").Preload("Order.Address").Where("id=?", id).Find(&orders)
	c.JSON(200, gin.H{
		"Product":          orders.Product.ProductName,
		"Amount":           orders.Amount,
		"Coupon":           orders.Order.Coupon.Code,
		"Status":           orders.Status,
		"Payment Method":   orders.Order.PaymentMethod,
		"Order Confirmed":  orders.CreatedAt,
		"Status Updated":   orders.UpdatedAt,
		"Quantity":         orders.Quantity,
		"Shipping Address": orders.Order.Address.ID,
	})
}

func CancelOrder(c *gin.Context) {
	id := c.Param("ID")
	var orderItem database.OrderItems
	if err := helper.DB.Preload("Order").Preload("Order.Coupon").Where("id=?", id).First(&orderItem); err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Can't Find the orderItem table.",
		})
		return
	}
	if orderItem.Status == "Cancelled" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "This Product alredy canclled ",
		})
		return
	}

	orderItem.Order.Amount -= orderItem.Amount
	if orderItem.Order.Amount <= float64(orderItem.Order.Coupon.Limit) && orderItem.Order.CouponID != 4 {
		orderItem.Order.Amount += orderItem.Order.Coupon.Amount
		//this showing an error the coupon id not update must check that
		orderItem.Order.CouponID = 4
		helper.DB.Model(&orderItem.Order).Updates(&orderItem.Order)
	}
	orderItem.Status = "Cancelled"
	helper.DB.Model(&orderItem.Order).Updates(&orderItem.Order)
	helper.DB.Save(&orderItem)
	c.JSON(200, "Order Cancelled.")

	// if {
	// 	orderItem.Order.Amount = orderItem.Order.Amount - (orderItem.Amount - orderItem.Order.Coupon.Amount)
	// 	fmt.Println(orderItem.Order.CouponID)
	// 	if err := helper.DB.Model(&orderItem.Order).Updates(&orderItem.Order); err.Error != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": "Can't Update the coupon id and amount.",
	// 		})
	// 		return
	// 	}
	// 	orderItem.Order.CouponID = 4
	// 	if err := helper.DB.Model(&orderItem.Order).Updates(&orderItem.Order); err.Error != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": "Can't Update the coupon id and amount.",
	// 		})
	// 		return
	// 	}
	// } else {
	// 	orderItem.Status = "Cancelled"
	// 	orderItem.Order.Amount -= orderItem.Amount
	// 	fmt.Println(orderItem.Order.Amount)
	// 	if err := helper.DB.Model(&orderItem.Order).Updates(&orderItem.Order); err.Error != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": "Can't Update the amount of order table.",
	// 		})
	// 		return
	// 	}
	// 	if err := helper.DB.Save(&orderItem); err.Error != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": "Can't save the orderItem table.",
	// 		})
	// 		return
	// 	}

	// }

}

func generateRandomNumber() string {
	const charset = "123456789"
	randomBytes := make([]byte, 8)
	_, err := rand.Read(randomBytes)
	if err != nil {
		fmt.Println(err)
	}
	for i, b := range randomBytes {
		randomBytes[i] = charset[b%byte(len(charset))]
	}
	return string(randomBytes)
}
