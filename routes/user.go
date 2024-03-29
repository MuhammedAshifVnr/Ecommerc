package routes

import (
	"ecom/controllers/users"
	"ecom/middleware"

	"github.com/gin-gonic/gin"
)

var UserRole = "user"

func UserRouters(r *gin.RouterGroup) {

	r.POST("/signup", users.Signup)
	r.POST("/login", users.Login)
	r.POST("/otp", users.OtpChecker)
	r.POST("/reotp", users.ResendOtp)

	r.POST("/forgot", users.EmailChecking)
	r.GET("/forgot", users.OtpValidation)
	r.PATCH("/forgot", users.UpdatePassword)

	r.GET("/home", users.Homepage)
	r.GET("/productDetail/:ID", users.ProductDetail)
	r.GET("/profile", middleware.AuthMiddleware(UserRole), users.Profile)
	r.PATCH("/profile", middleware.AuthMiddleware(UserRole), users.EditeProfile)

	r.GET("/address", middleware.AuthMiddleware(UserRole), users.Address)
	r.POST("/address", middleware.AuthMiddleware(UserRole), users.AddAddress)
	r.PUT("/address/:ID", middleware.AuthMiddleware(UserRole), users.AddressEdit)
	r.DELETE("/address/:ID", middleware.AuthMiddleware(UserRole), users.AddressDelete)

	r.GET("/cart", middleware.AuthMiddleware(UserRole), users.Cart)
	r.POST("/cart/:ID", middleware.AuthMiddleware(UserRole), users.AddCart)
	r.PATCH("/cart/:ID", middleware.AuthMiddleware(UserRole), users.CartQuantity)
	r.DELETE("/cart/:ID", middleware.AuthMiddleware(UserRole), users.CartDelete)
	r.POST("/checkout", middleware.AuthMiddleware(UserRole), users.Testcheckout)
	r.POST("/review/:ID", middleware.AuthMiddleware(UserRole), users.CreatReview)

	r.GET("/order", middleware.AuthMiddleware(UserRole), users.Order)
	r.GET("/order-item/:ID", middleware.AuthMiddleware(UserRole), users.OrderDetils)
	r.PATCH("/order/:ID", middleware.AuthMiddleware(UserRole), users.CancelOrder)

	r.GET("/whislist", middleware.AuthMiddleware(UserRole), users.Whislist)
	r.POST("/whislist/:ID", middleware.AuthMiddleware(UserRole), users.AddWhislist)
	r.DELETE("/whislist/:ID", middleware.AuthMiddleware(UserRole), users.DeleteWhislist)

	r.GET("/payment", func(c *gin.Context) {
		c.HTML(200, "payment.html", nil)
	})
	r.POST("/razorpay-payment", users.HandleRazorpayPayment)

	r.GET("/search-product", users.SeaechProduct)

	r.GET("/logout", users.Logout)


	r.GET("/invoice/:ID",middleware.AuthMiddleware(UserRole),users.GenerateInvoice)
}
