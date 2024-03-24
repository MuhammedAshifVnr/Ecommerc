package routes

import (
	"ecom/jwt"
	"ecom/controllers/users"

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
	r.GET("/profile", jwt.AuthMiddleware(UserRole), users.Profile)
	r.PATCH("/profile", jwt.AuthMiddleware(UserRole), users.EditeProfile)

	r.GET("/address", jwt.AuthMiddleware(UserRole), users.Address)
	r.POST("/address", jwt.AuthMiddleware(UserRole), users.AddAddress)
	r.PUT("/address/:ID", jwt.AuthMiddleware(UserRole), users.AddressEdit)
	r.DELETE("/address/:ID", jwt.AuthMiddleware(UserRole), users.AddressDelete)

	r.GET("/cart", jwt.AuthMiddleware(UserRole), users.Cart)
	r.POST("/cart/:ID", jwt.AuthMiddleware(UserRole), users.AddCart)
	r.PATCH("/cart/:ID", jwt.AuthMiddleware(UserRole), users.CartQuantity)
	r.DELETE("/cart/:ID", jwt.AuthMiddleware(UserRole), users.CartDelete)
	r.POST("/checkout", jwt.AuthMiddleware(UserRole), users.Testcheckout)
	r.POST("/review/:ID", jwt.AuthMiddleware(UserRole), users.CreatReview)

	r.GET("/order", jwt.AuthMiddleware(UserRole), users.Order)
	r.GET("/order-item/:ID", jwt.AuthMiddleware(UserRole), users.OrderDetils)
	r.PATCH("/order/:ID", jwt.AuthMiddleware(UserRole), users.CancelOrder)

	r.GET("/whislist", jwt.AuthMiddleware(UserRole), users.Whislist)
	r.POST("/whislist/:ID",jwt.AuthMiddleware(UserRole),users.AddWhislist)
	r.DELETE("/whislist/:ID",jwt.AuthMiddleware(UserRole),users.DeleteWhislist)

	r.GET("/payment",func (c *gin.Context)  {
		c.HTML(200,"payment.html",nil)
	})
	r.POST("/razorpay-payment",users.HandleRazorpayPayment)
	

	r.GET("/search-product", jwt.AuthMiddleware(UserRole), users.SeaechProduct)

	r.GET("/logout", users.Logout)
}
