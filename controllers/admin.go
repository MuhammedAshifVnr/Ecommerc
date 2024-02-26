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
	r.POST("/addproduct", admin.AddProduct)
	r.PATCH("/edit/:ID", admin.EditProdect)
	r.PATCH("/imagedit/:ID",admin.ImageEditing)
	r.DELETE("/delete/:ID", admin.Delete)
	//users
	r.GET("/users",admin.UsersList)
	r.PATCH("/userblock/:ID",admin.UserStatus)
//category
	r.GET("/category",admin.Category)
	r.POST("/addcategory",admin.AddCategory)
	r.PUT("/editcategory/:ID",admin.EditCategory)
	r.PATCH("/blockcategory/:ID",admin.BlockCategory)
	r.DELETE("/deletecategory/:ID",admin.DeleteCategory)

//helper
	r.PATCH("/recover/:ID",admin.DeleteRecovery)
//image adding
	r.POST("/imageadding",admin.ImageAdding)

}
