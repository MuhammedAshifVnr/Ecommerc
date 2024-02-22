package controllers

import (
	"ecom/routers/users"

	"github.com/gin-gonic/gin"
)

func UserRouters(r *gin.RouterGroup) {

	r.POST("/signup", users.Signup)
	r.POST("/login", users.Login)
	r.POST("/otp", users.OtpChecker)
	r.POST("/reotp",users.ResendOtp)

}
