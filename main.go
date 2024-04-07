package main

import (
	_ "ecom/docs"
	"ecom/helper"
	"ecom/routes"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	helper.EnvLoader()
	helper.DbConnect()
}


// @title ECOM
// @version 1.0
// @description This is a sample Gin API with Swagger documentation.
// @host localhost:8080
// @BasePath /
func main() {
	router := gin.Default()
	router.LoadHTMLGlob("temp/*")

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	User := router.Group("/user")
	routes.UserRouters(User)

	Admin := router.Group("/admin")
	routes.AdminRouters(Admin)

	router.Run(":8080")

}
