package main

import (
	"ecom/helper"
	"ecom/routes"

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
	router.LoadHTMLGlob("temp/*")
	
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	User := router.Group("/user")
	routes.UserRouters(User)

	Admin := router.Group("/admin")
	routes.AdminRouters(Admin)

	router.Run(":8080")

}
