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

	r.GET("/address", users.Address)
	r.POST("/address", users.AddAddress)
	r.PUT("/address/:ID", users.AddressEdit)
	r.DELETE("/address/:ID", users.AddressDelete)

	r.GET("/logout",users.Logout)
}
