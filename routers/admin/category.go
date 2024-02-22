package admin

import (
	"ecom/database"
	"ecom/helper"

	"github.com/gin-gonic/gin"
)

func Category(c *gin.Context) {
	var find []database.Category

	helper.DB.Find(&find)
	for _, v := range find {
		c.JSON(200, gin.H{
			"Name":        v.Name,
			"Description": v.Description,
			"Status":      v.Status,
		})
	}
}

func AddCategory(c *gin.Context) {
	var cat database.Category

	c.ShouldBindJSON(&cat)

	if err := helper.DB.Create(&cat); err.Error != nil {
		c.JSON(404, "Duplicate found.")
		return
	}
	c.JSON(200, "Successfully added.")

}

func EditCategory(c *gin.Context) {
	var find database.Category

	id := c.Param("ID")

	c.ShouldBindJSON(&find)
	if err := helper.DB.Where("id=?", id).Updates(&find); err.Error != nil {
		c.JSON(404, "error")
	}
	c.JSON(200, "Successfully Edited.")
}

func BlockCategory(c *gin.Context) {
	var find database.Category
	var prod []database.Product

	id := c.Param("ID")

	helper.DB.Where("id=?", id).First(&find)
	if find.Status == "Active" {
		helper.DB.Model(&find).Update("Status", "Block")
		helper.DB.Where("category_id=?", id).Find(&prod)
		for _, v := range prod {
			helper.DB.Model(&v).Update("status", "Block")
		}
		c.JSON(200, "Successfully Blocked.")
	} else {
		helper.DB.Model(&find).Update("Status", "Active")
		helper.DB.Where("category_id=?", id).Find(&prod)
		for _, v := range prod {
			helper.DB.Model(&v).Update("status", "Active")
		}
		c.JSON(200, "Successfully Actived.")
	}
}

func DeleteCategory(c *gin.Context) {
	var find database.Category

	id := c.Param("ID")

	helper.DB.Where("id=?", id).Delete(&find)
	c.JSON(200, "Category deleted Successfully.")
}
