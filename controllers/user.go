package controllers

import (
	"ecom/routers/users"

	"github.com/gin-gonic/gin"
)

func UserRouters(r *gin.RouterGroup) {

	r.POST("/signup", users.Signup)
	r.POST("/login", users.Login)
	r.POST("/otp", users.OtpChecker)
	r.POST("/reotp", users.ResendOtp)
	r.GET("/home", users.Homepage)
	r.GET("/productDetail/:ID", users.ProductDetail)
	r.GET("/profile", users.Profile)

	r.POST("/forgot",users.EmailChecking)
	r.GET("/forgot",users.OtpValidation)
	r.PATCH("/forgot",users.UpdatePassword)

	r.GET("/address", users.Address)
	r.POST("/address", users.AddAddress)
	r.PUT("/address/:ID", users.AddressEdit)
	r.DELETE("/address/:ID", users.AddressDelete)

	r.GET("/cart",users.Cart)
	r.POST("/cart/:ID",users.AddCart)
	r.PATCH("/cart/:ID",users.CartQuantity)
	r.DELETE("/cart/:ID",users.CartDelete)
	r.POST("/checkout",users.CheckOut)
	r.POST("/review/:ID",users.CreatReview)

	r.GET("/order",users.Order)
	r.GET("/order-item/:ID",users.OrderDetils)
	r.PATCH("/order/:ID",users.CancelOrder)

	r.GET("/search-product",users.SeaechProduct)

	r.GET("/logout",users.Logout)
}
