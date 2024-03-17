package users

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/razorpay/razorpay-go"
)

var key = os.Getenv("KEYID")
var secret = os.Getenv("KEYSECRET")
var razorpayClient = razorpay.NewClient(key, secret)

func OrderCreat(c *gin.Context) {
	paymentOrder, err := razorpayClient.Order.Create(map[string]interface{}{
		"amount":   1554,
		"currency": "INR",
		// ... other payment parameters
	}, nil)

	if err != nil {
		// Handle the error
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Return the payment order details to the client
	c.JSON(http.StatusOK, gin.H{
		"order_id": paymentOrder,
		// ... other payment order details
	})
}
