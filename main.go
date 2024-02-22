package main

import (
	"ecom/controllers"
	"ecom/helper"

	"github.com/gin-gonic/gin"
)

func init() {
	helper.USerOtp = make(map[string]string)
	helper.EnvLoader()
	helper.DbConnect()
}

func main() {

	router := gin.Default()

	User := router.Group("/user")
	controllers.UserRouters(User)

	Admin := router.Group("/admin")
	controllers.AdminRouters(Admin)

	router.Run(":8080")
	

}
