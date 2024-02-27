package controllers

import (
	"ecom/routers/admin"

	"github.com/gin-gonic/gin"
)

func AdminRouters(r *gin.RouterGroup) {
//login
	r.POST("/login", admin.AdminLogin)
	r.GET("/home", admin.HomePage)

//product
	r.GET("/product", admin.Product)
	r.POST("/product", admin.AddProduct)
	r.PUT("/product/:ID", admin.EditProdect)
	r.DELETE("/product/:ID", admin.Delete)
	
	//users
	r.GET("/users",admin.UsersList)
	r.PATCH("/users/:ID",admin.UserStatus)

	//category
	r.GET("/category",admin.Category)
	r.POST("/category",admin.AddCategory)
	r.PUT("/category/:ID",admin.EditCategory)
	r.PATCH("/category/:ID",admin.BlockCategory)
	r.DELETE("/category/:ID",admin.DeleteCategory)

	//rating
	r.POST("/rating",admin.RatingAdding)
//helper
	r.PATCH("/recover/:ID",admin.DeleteRecovery)
//image adding
	//r.POST("/imageadding",admin.ImageAdding)

}
