package users

import (
	"ecom/database"
	"ecom/helper"
	"fmt"

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
			"ID":v.ID,
		})
	}
}

func ProductDetail(c *gin.Context) {
	var find database.Product
	var stock string
	var table []database.Product
	var rating database.Ratings

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
		"Images":      find.ImageUrls,
	})
	helper.DB.Where("product_id=?", id).First(&rating)
	result := rating.Rating / float32(rating.Users)
	fmt.Println(result)
	c.JSON(200, result)
	c.JSON(200, "Recommend Products")
	for i := 0; i < len(table); i++ {
		if find.ID != table[i].ID {
			c.JSON(200, gin.H{
				"Image": table[i].ImageUrls,
				"Name":  table[i].ProductName,
				"Prize": table[i].ProductPrize,
			})
		}
	}
}
