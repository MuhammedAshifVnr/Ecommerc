package main

import (
	"ecom/controllers"
	"ecom/helper"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func init() {
	helper.EnvLoader()
	helper.DbConnect()
}

func main() {
	router := gin.Default()
	
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	User := router.Group("/user")
	controllers.UserRouters(User)

	Admin := router.Group("/admin")
	controllers.AdminRouters(Admin)

	router.Run(":8080")

}
