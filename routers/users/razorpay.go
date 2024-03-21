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

	"github.com/gin-gonic/gin"
	"github.com/razorpay/razorpay-go"
)

var razorpayClient = razorpay.NewClient("rzp_test_ZW342DjmlVWLMO", "H8uZz3CTyagwZqO53aorSUj6")

func HandleRazorpayPayment(c *gin.Context) {
	var respons map[string]string
	if err := c.ShouldBindJSON(&respons); err != nil {
		fmt.Println("error:", err)
		return
	}
	err:=RazorPaymentVerification(respons["razorpay_signature"],respons["razorpay_order_id"],respons["razorpay_payment_id"])
	if err!=nil{
		fmt.Println("eroooooor:",err)
		return
	}else{
		fmt.Println("Payment Done.")
	}
	fmt.Println(respons)
	payment:=database.Transactions{
		PaymentID: respons["razorpay_payment_id"],
		OrderID: respons["razorpay_order_id"],
	}
	helper.DB.Create(&payment)
	c.JSON(http.StatusOK, gin.H{"message": "Payment response received successfully"})
}

func RazorPaymentVerification(sign, orderId, paymentId string) error {
	signature := sign
	secret := "H8uZz3CTyagwZqO53aorSUj6"
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
