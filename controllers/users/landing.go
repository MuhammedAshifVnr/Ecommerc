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
		var rating []database.Review
		helper.DB.Where("product_id=?", v.ID).Find(&rating)
		avg := AvgRating(rating)
		discount := ProductOffer(v.ID)
		c.JSON(200, gin.H{
			"Name":     v.ProductName,
			"Prize":    v.ProductPrice - discount,
			"Category": v.Category.Name,
			"Rating":   avg,
			"ID":       v.ID,
		})
	}
}

func ProductDetail(c *gin.Context) {
	var find database.Product
	var stock string
	var table []database.Product
	var review database.Review

	id := c.Param("ID")
	discount := ProductOffer(id)
	helper.DB.Preload("Category").First(&find, "id=?", id)
	if find.Quantity == 0 {
		stock = "Out of Stock"
	} else if find.Quantity <= 5 {
		stock = "Only few items"
	} else {
		stock = "Available"
	}
	helper.DB.Where("category_id=?", find.CategoryId).Find(&table)
	var rating []database.Review
	helper.DB.Where("product_id=?", find.ID).Find(&rating)
	avg := AvgRating(rating)
	c.JSON(200, gin.H{
		"Name":        find.ProductName,
		"Prize":       find.ProductPrice - discount,
		"Stock":       stock,
		"Size":        find.Size,
		"Description": find.Description,
		"Category":    find.Category.Name,
		"Images":      find.ImageUrls,
		"Discount":    discount,
		"Rating":      avg,
	})
	for _, v := range rating {
		c.JSON(200, gin.H{
			"Rating": v.Rating,
			"Review": v.Comment,
		})
	}
	helper.DB.Where("product_id=?", id).First(&review)
	c.JSON(200, "Recommend Products")
	for i := 0; i < len(table); i++ {
		if find.ID != table[i].ID {
			c.JSON(200, gin.H{
				"Image": table[i].ImageUrls,
				"Name":  table[i].ProductName,
				"Prize": table[i].ProductPrice,
			})
		}
	}
}

func AvgRating(ratings []database.Review) float64 {
	avg := 0.0
	if len(ratings) == 0 {
		return avg
	}
	for _, v := range ratings {
		avg += v.Rating
	}
	return avg / float64(len(ratings))
}
