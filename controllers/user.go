package controllers

import (
	"ecom/middleware"
	"ecom/routers/users"

	"github.com/gin-gonic/gin"
)

var roleUser = "user"

func UserRouters(r *gin.RouterGroup) {

	r.POST("/signup", users.Signup)
	r.POST("/login", users.Login)
	r.POST("/otp", users.OtpChecker)
	r.POST("/reotp", users.ResendOtp)

	r.POST("/forgot", users.EmailChecking)
	r.GET("/forgot", users.OtpValidation)
	r.PATCH("/forgot", users.UpdatePassword)

	r.GET("/home", middleware.AuthMiddleware(roleUser), users.Homepage)
	r.GET("/productDetail/:ID", middleware.AuthMiddleware(roleUser), users.ProductDetail)
	r.GET("/profile", middleware.AuthMiddleware(roleUser), users.Profile)
	r.PATCH("/profile", middleware.AuthMiddleware(roleUser), users.EditeProfile)

	r.GET("/address", middleware.AuthMiddleware(roleUser), users.Address)
	r.POST("/address", middleware.AuthMiddleware(roleUser), users.AddAddress)
	r.PUT("/address/:ID", middleware.AuthMiddleware(roleUser), users.AddressEdit)
	r.DELETE("/address/:ID", middleware.AuthMiddleware(roleUser), users.AddressDelete)

	r.GET("/cart", middleware.AuthMiddleware(roleUser), users.Cart)
	r.POST("/cart/:ID", middleware.AuthMiddleware(roleUser), users.AddCart)
	r.PATCH("/cart/:ID", middleware.AuthMiddleware(roleUser), users.CartQuantity)
	r.DELETE("/cart/:ID", middleware.AuthMiddleware(roleUser), users.CartDelete)
	r.POST("/checkout", middleware.AuthMiddleware(roleUser), users.CheckOut)
	r.POST("/review/:ID", middleware.AuthMiddleware(roleUser), users.CreatReview)

	r.GET("/order", middleware.AuthMiddleware(roleUser), users.Order)
	r.GET("/order-item/:ID", middleware.AuthMiddleware(roleUser), users.OrderDetils)
	r.PATCH("/order/:ID", middleware.AuthMiddleware(roleUser), users.CancelOrder)

	r.GET("/search-product", middleware.AuthMiddleware(roleUser), users.SeaechProduct)

	r.GET("/logout", users.Logout)
}
