package users

import (
	"ecom/database"
	"ecom/helper"

	"github.com/gin-gonic/gin"
)

func Homepage(c *gin.Context) {
	var find []database.Product

	helper.DB.Preload("Category").Find(&find)

	for _, v := range find {
		c.JSON(200, gin.H{
			"Name":     v.ProductName,
			"Prize":    v.ProductPrize,
			"Category": v.Category.Name,
		})
	}
}

func ProductDetail(c *gin.Context) {
	var find database.Product
	var stock string
	var table []database.Product

	id := c.Param("ID")

	helper.DB.Preload("Category").First(&find, "id=?", id)
	if find.Quantity == 0 {
		stock = "Out of Stock"
	} else if find.Quantity <= 5 {
		stock = "Only few items"
	} else {
		stock = "Available"
	}
	helper.DB.Where("category_id=?", find.CategoryId).Find(&table)
	c.JSON(200, gin.H{
		"Name":        find.ProductName,
		"Prize":       find.ProductPrize,
		"Stock":       stock,
		"Size":        find.Size,
		"Description": find.Description,
		"Category":    find.Category.Name,
		"Image_1":     find.ImageUrl1,
		"Image_2":     find.ImageUrl2,
		"Image_3":     find.ImageUrl3,
	})
c.JSON(200,"Recommend Products")
	for i := 0; i < len(table); i++ {
		if find.ID != table[i].ID {
			c.JSON(200, gin.H{
				"Image": table[i].ImageUrl1,
				"Name":  table[i].ProductName,
				"Prize": table[i].ProductPrize,
			})
		}
	}
}
