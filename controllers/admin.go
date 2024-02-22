package controllers

import (
	"ecom/routers/admin"

	"github.com/gin-gonic/gin"
)

func AdminRouters(r *gin.RouterGroup) {

	r.POST("/login", admin.AdminLogin)
	r.GET("/home", admin.HomePage)

	r.GET("/product", admin.Product)
	r.POST("/addproduct", admin.AddProduct)
	r.PUT("/edit/:ID", admin.EditProdect)
	r.DELETE("/delete/:ID", admin.Delete)
	
	r.GET("/users",admin.UsersList)
	r.PATCH("/userblock/:ID",admin.UserStatus)

	r.GET("/category",admin.Category)
	r.POST("/addcategory",admin.AddCategory)
	r.PUT("/editcategory/:ID",admin.EditCategory)
	r.PATCH("/blockcategory/:ID",admin.BlockCategory)
	r.DELETE("/deletecategory/:ID",admin.DeleteCategory)


	r.PATCH("/recover/:ID",admin.DeleteRecovery)

	r.POST("/imageadding",admin.ImageAdding)

}
