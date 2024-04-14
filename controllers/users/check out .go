package users

import (
	"crypto/rand"
	"ecom/database"
	"ecom/helper"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/razorpay/razorpay-go"
)

// @Summary Checkout Page
// @Description Checkout Page
// @Tags User-Cart
// @Accept	multipart/form-data
// @Produce json
// @Param payment formData string true "Payment Method"
// @Param address formData string true "Address"
// @Param coupon formData string false "coupon"
// @Router /user/checkout [post]
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
		amount := (item.Product.ProductPrice - ProductOffer(item.ProductID)) * float64(item.Quantity)
		if item.Quantity >= uint(item.Product.Quantity) {
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
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "You can't use this coupon."})
		return
	}
	totalAmount -= coupon.Amount

	if totalAmount <= 1500 {
		totalAmount += 40
		tx.Model(&database.Order{}).Where("id=?", orderID).Update("dlivery_charge", 40)
	}

	tx.Model(&database.Order{}).Where("id=?", orderID).Update("amount", totalAmount)

	if order.PaymentMethod == "ONLINE" {
		RazorpayClient := razorpay.NewClient(os.Getenv("RAZORPAY_ID"), os.Getenv("RAZORPAY_SECRET_ID"))
		paymentOrder, err := RazorpayClient.Order.Create(map[string]interface{}{
			"amount":   totalAmount * 100,
			"currency": "INR",
			"receipt":  strconv.Itoa(int(orderID)),
		}, nil)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{
				"error===": err.Error(),
			})
			return
		}
		payment := database.Transactions{
			OrderID: paymentOrder["id"].(string),
			Amount:  totalAmount,
			Status:  "Failed",
			Receipt: orderID,
		}
		helper.DB.Create(&payment)
		tx.Commit()
		helper.DB.Model(&database.Order{}).Where("id=?", orderID).Update("payment_id", paymentOrder["id"])
		helper.DB.Where("user_id=?", userId).Delete(&database.Cart{})
		c.JSON(http.StatusOK, gin.H{
			"message":   "Order Placed Successfully.",
			"Amount":    totalAmount,
			"PaymentID": paymentOrder["id"],
		})
	} else if order.PaymentMethod == "COD" {
		if totalAmount >= 1000 {
			tx.Rollback()
			c.JSON(404, gin.H{
				"Error":   "COD not available for this order",
				"Massege": "Choose any another payment menthod.",
			})
			return
		}
		tx.Commit()
		helper.DB.Where("user_id=?", userId).Delete(&database.Cart{})
		c.JSON(http.StatusOK, gin.H{
			"message": "Order Placed Successfully. Product Get Soon",
			"Amount":  totalAmount,
		})
	} else if order.PaymentMethod == "WALLET" {
		var wallet database.Wallet
		helper.DB.Where("user_id=?", userId).First(&wallet)
		if totalAmount <= wallet.Amount {
			wallet.Amount -= totalAmount
			helper.DB.Save(&wallet)
			tx.Commit()
			helper.DB.Where("user_id=?", userId).Delete(&database.Cart{})
			c.JSON(http.StatusOK, gin.H{
				"message": "Order Placed Successfully.",
				"Amount":  totalAmount,
			})
		} else {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Inefficient Balance",
			})
		}

	}

}

// @Summary Order Listing
// @Description Order Listing
// @Tags User-Order
// @Accept  json
// @Produce  json
// @Router /user/order [get]
func Order(c *gin.Context) {
	var orders []database.Order
	userID := c.GetUint("userID")
	helper.DB.Preload("Coupon").Where("user_id=?", userID).Order("created_at DESC").Find(&orders)
	var order_list []gin.H
	for _, v := range orders {
		order_list = append(order_list, gin.H{
			"ID":            v.ID,
			"created":       v.CreatedAt,
			"paymentMethod": v.PaymentMethod,
			"coupon":        v.Coupon.Code,
			"amount":        v.Amount,
		})
	}

	c.JSON(200, gin.H{"code": 200, "status": "Success", "data": gin.H{"orders": order_list}})
}

// @Summary Order item listing
// @Descripton Order item listing
// @Tags User-Order
// @Accept  json
// @Produce  json
// @Param ID path string true "Order ID"
// @Router /user/order-item/{ID} [get]
func OrderDetils(c *gin.Context) {
	var orders database.OrderItems
	id := c.Param("ID")
	helper.DB.Preload("Product").Preload("Order.Coupon").Preload("Order.Address").Where("order_id=?", id).Find(&orders)
	item_list := gin.H{
		"id":           orders.ID,
		"amount":       orders.Amount,
		"productName":  orders.Product.ProductName,
		"productImage": orders.Product.ImageUrls,
		"orderID":      orders.OrderID,
		"status":       orders.Status,
		"quantity":     orders.Quantity,
	}

	c.JSON(200, gin.H{"code": 200, "status": "Success", "data": gin.H{"Items": item_list}})
}

// @Summary Order Cancelation
// @Description Order Cancelation
// @Tags User-Order
// @Accept  json
// @Produce  json
// @Param ID path string true "Order ID"
// @Router /user/order/{ID} [patch]
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
		if orderItem.Order.PaymentMethod == "ONLINE" {
			var wellet database.Wallet
			if err := helper.DB.Where("user_id=?", c.GetUint("userID")).First(&wellet); err.Error != nil {
				fmt.Println("error is =", err.Error)
				return
			}
			fmt.Println(wellet)
			wellet.Amount += orderItem.Amount - orderItem.Order.Coupon.Amount
			helper.DB.Save(&wellet)
			orderItem.Status = "Cancelled"
			helper.DB.Save(&orderItem)
		}
	} else {
		if orderItem.Order.PaymentMethod == "ONLINE" {
			var wellet database.Wallet
			if err := helper.DB.Where("user_id=?", c.GetUint("userID")).First(&wellet); err.Error != nil {
				fmt.Println("error is =", err.Error)
				return
			}
			fmt.Println(wellet)
			wellet.Amount += orderItem.Amount
			helper.DB.Save(&wellet)
		}
		orderItem.Status = "Cancelled"
		helper.DB.Model(&orderItem.Order).Updates(&orderItem.Order)
		helper.DB.Save(&orderItem)
	}
	c.JSON(200, gin.H{"Massage": "Order Cancelled."})

}

func generateRandomNumber() string {
	const charset = "123456789"
	randomBytes := make([]byte, 6)
	_, err := rand.Read(randomBytes)
	if err != nil {
		fmt.Println(err)
	}
	for i, b := range randomBytes {
		randomBytes[i] = charset[b%byte(len(charset))]
	}
	return string(randomBytes)
}
