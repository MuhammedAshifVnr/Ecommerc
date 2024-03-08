package controllers

import (
	"ecom/middleware"
	"ecom/routers/admin"
	"ecom/routers/users"

	"github.com/gin-gonic/gin"
)

var role = "admin"

func AdminRouters(r *gin.RouterGroup) {
	//login
	r.POST("/login", admin.AdminLogin)
	r.GET("/logout",admin.Logout)
	r.GET("/home", middleware.AuthMiddleware(role), admin.HomePage)

	//product
	r.GET("/product", middleware.AuthMiddleware(role), admin.Product)
	r.POST("/product", middleware.AuthMiddleware(role), admin.AddProduct)
	r.PUT("/product/:ID", middleware.AuthMiddleware(role), admin.EditProdect)
	r.DELETE("/product/:ID", middleware.AuthMiddleware(role), admin.Delete)

	//users
	r.GET("/users", middleware.AuthMiddleware(role), admin.UsersList)
	r.PATCH("/users/:ID", middleware.AuthMiddleware(role), admin.UserStatus)

	//category
	r.GET("/category", middleware.AuthMiddleware(role), admin.Category)
	r.POST("/category", middleware.AuthMiddleware(role), admin.AddCategory)
	r.PUT("/category/:ID", middleware.AuthMiddleware(role), admin.EditCategory)
	r.PATCH("/category/:ID", middleware.AuthMiddleware(role), admin.BlockCategory)
	r.DELETE("/category/:ID", middleware.AuthMiddleware(role), admin.DeleteCategory)

	//coupons
	r.GET("/coupon", middleware.AuthMiddleware(role), admin.Coupons)
	r.POST("/coupon", middleware.AuthMiddleware(role), admin.AddCoupons)
	r.DELETE("/coupon/:ID", middleware.AuthMiddleware(role), admin.DeleteCoupon)

	r.GET("/order", middleware.AuthMiddleware(role), admin.Orders)
	r.PATCH("/order/update/:ID",middleware.AuthMiddleware(role),admin.UpdateOrder)
	r.PATCH("/order/:ID", middleware.AuthMiddleware(role), users.CancelOrder)

	//helper
	r.PATCH("/recover/:ID", admin.DeleteRecovery)


}
