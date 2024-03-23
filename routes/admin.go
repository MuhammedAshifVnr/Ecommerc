package routes

import (
	"ecom/jwt"
	"ecom/routers/admin"
	"ecom/routers/users"

	"github.com/gin-gonic/gin"
)

var AdminRole = "admin"

func AdminRouters(r *gin.RouterGroup) {
	//login
	r.POST("/login", admin.AdminLogin)
	r.GET("/logout", admin.Logout)
	r.GET("/home",jwt.AuthMiddleware(AdminRole), admin.HomePage)

	//product
	r.GET("/product",jwt.AuthMiddleware(AdminRole), admin.Product)
	r.POST("/product",jwt.AuthMiddleware(AdminRole), admin.AddProduct)
	r.PUT("/product/:ID",jwt.AuthMiddleware(AdminRole), admin.EditProdect)
	r.DELETE("/product/:ID",jwt.AuthMiddleware(AdminRole), admin.Delete)

	//users
	r.GET("/users",jwt.AuthMiddleware(AdminRole), admin.UsersList)
	r.PATCH("/users/:ID",jwt.AuthMiddleware(AdminRole), admin.UserStatus)

	//category
	r.GET("/category",jwt.AuthMiddleware(AdminRole), admin.Category)
	r.POST("/category",jwt.AuthMiddleware(AdminRole), admin.AddCategory)
	r.PUT("/category/:ID",jwt.AuthMiddleware(AdminRole), admin.EditCategory)
	r.PATCH("/category/:ID",jwt.AuthMiddleware(AdminRole), admin.BlockCategory)
	r.DELETE("/category/:ID", jwt.AuthMiddleware(AdminRole),admin.DeleteCategory)

	//coupons
	r.GET("/coupon",jwt.AuthMiddleware(AdminRole), admin.Coupons)
	r.POST("/coupon",jwt.AuthMiddleware(AdminRole), admin.AddCoupons)
	r.DELETE("/coupon/:ID",jwt.AuthMiddleware(AdminRole), admin.DeleteCoupon)

	r.GET("/order",jwt.AuthMiddleware(AdminRole), admin.Orders)
	r.PATCH("/order/update/:ID",jwt.AuthMiddleware(AdminRole), admin.UpdateOrder)
	r.PATCH("/order/:ID",jwt.AuthMiddleware(AdminRole), users.CancelOrder)

	//helper
	r.PATCH("/recover/:ID", admin.DeleteRecovery)

}
