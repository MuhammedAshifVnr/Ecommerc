package routes

import (
	"ecom/controllers/admin"
	"ecom/controllers/users"
	"ecom/middleware"

	"github.com/gin-gonic/gin"
)

var AdminRole = "admin"

func AdminRouters(r *gin.RouterGroup) {
	//login
	r.POST("/login", admin.AdminLogin)
	r.GET("/logout", admin.Logout)
	r.GET("/home", middleware.AuthMiddleware(AdminRole), admin.HomePage)

	//product
	r.GET("/product", middleware.AuthMiddleware(AdminRole), admin.Product)
	r.POST("/product", middleware.AuthMiddleware(AdminRole), admin.AddProduct)
	r.PUT("/product/:ID", middleware.AuthMiddleware(AdminRole), admin.EditProdect)
	r.DELETE("/product/:ID", middleware.AuthMiddleware(AdminRole), admin.Delete)

	//users
	r.GET("/users", middleware.AuthMiddleware(AdminRole), admin.UsersList)
	r.PATCH("/users/:ID", middleware.AuthMiddleware(AdminRole), admin.UserStatus)

	//category
	r.GET("/category", middleware.AuthMiddleware(AdminRole), admin.Category)
	r.POST("/category", middleware.AuthMiddleware(AdminRole), admin.AddCategory)
	r.PUT("/category/:ID", middleware.AuthMiddleware(AdminRole), admin.EditCategory)
	r.PATCH("/category/:ID", middleware.AuthMiddleware(AdminRole), admin.BlockCategory)
	r.DELETE("/category/:ID", middleware.AuthMiddleware(AdminRole), admin.DeleteCategory)

	//coupons
	r.GET("/coupon", middleware.AuthMiddleware(AdminRole), admin.Coupons)
	r.POST("/coupon", middleware.AuthMiddleware(AdminRole), admin.AddCoupons)
	r.DELETE("/coupon/:ID", middleware.AuthMiddleware(AdminRole), admin.DeleteCoupon)

	r.GET("/order", middleware.AuthMiddleware(AdminRole), admin.Orders)
	r.PATCH("/order/update/:ID", middleware.AuthMiddleware(AdminRole), admin.UpdateOrder)
	r.PATCH("/order/:ID", middleware.AuthMiddleware(AdminRole), users.CancelOrder)

	r.GET("/sort",admin.ProductSorting)

	r.GET("/salesreport",admin.DownloadReport)
	//helper
	r.PATCH("/recover/:ID", admin.DeleteRecovery)

}
