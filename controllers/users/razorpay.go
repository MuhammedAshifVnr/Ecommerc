package users

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"ecom/database"
	"ecom/helper"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func HandleRazorpayPayment(c *gin.Context) {
	var respons map[string]string
	if err := c.ShouldBindJSON(&respons); err != nil {
		fmt.Println("error:", err)
		return
	}
	err := RazorPaymentVerification(respons["razorpay_signature"], respons["razorpay_order_id"], respons["razorpay_payment_id"])
	if err != nil {
		fmt.Println("eroooooor:", err)
		return
	} else {
		fmt.Println("Payment Done.")
	}
	fmt.Println(respons)
	payment := database.Transactions{
		PaymentID: respons["razorpay_payment_id"],
		Status:    "Success",
	}
	helper.DB.Where("order_id=?", respons["razorpay_order_id"]).Updates(&payment)
	c.JSON(http.StatusOK, gin.H{"message": "Payment response received successfully"})
}

func RazorPaymentVerification(sign, orderId, paymentId string) error {
	signature := sign
	secret := os.Getenv("RAZORPAY_SECRET_ID")
	data := orderId + "|" + paymentId
	h := hmac.New(sha256.New, []byte(secret))
	_, err := h.Write([]byte(data))
	if err != nil {
		panic(err)
	}
	sha := hex.EncodeToString(h.Sum(nil))
	if subtle.ConstantTimeCompare([]byte(sha), []byte(signature)) != 1 {
		return errors.New("PAYMENT FAILED")
	} else {
		return nil
	}
}
