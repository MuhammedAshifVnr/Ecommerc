package users

import (
	"ecom/database"
	"ecom/helper"
	"math"

	"github.com/gin-gonic/gin"
)

// @Summary      Get All Products
// @Description  Get All Products
// @Tags         User-Product
// @Produce      json
// @Router       /user/home [get]
func Homepage(c *gin.Context) {
	var find []database.Product

	helper.DB.Preload("Category").Preload("Offers").Find(&find)
	var product_list []gin.H
	for _, v := range find {
		var rating []database.Review
		helper.DB.Where("product_id=?", v.ID).Find(&rating)
		avg := AvgRating(rating)
		discount := ProductOffer(v.ID)

		product_list = append(product_list, gin.H{
			"Name":     v.ProductName,
			"Prize":    v.ProductPrice - discount,
			"Category": v.Category.Name,
			"Rating":   avg,
			"ID":       v.ID,
		})
	}
	c.JSON(200, gin.H{
		"code":     200,
		"status":   "Success",
		"products": product_list,
	})
}

// @Summary Product Detail
// @Description Product Detail
// @Tags User-Product
// @Produce json
// @Param ID path string true "Product ID"
// @Router /user/productDetail/{ID} [get]
func ProductDetail(c *gin.Context) {
	var find database.Product
	var stock string
	var table []database.Product

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
	product := gin.H{
		"Name":        find.ProductName,
		"Prize":       find.ProductPrice - discount,
		"Stock":       stock,
		"Size":        find.Size,
		"Description": find.Description,
		"Category":    find.Category.Name,
		"Images":      find.ImageUrls,
		"Discount":    discount,
		"Rating":      avg,
	}
	var ratings []gin.H
	for _, v := range rating {
		ratings = append(ratings, gin.H{
			"Rating": v.Rating,
			"Review": v.Comment,
		})
	}

	var recomend []gin.H
	for _, v := range table {
		if find.ID != v.ID {
			recomend = append(recomend, gin.H{
				"Image": v.ImageUrls,
				"Name":  v.ProductName,
				"Prize": v.ProductPrice,
			})
		}
	}
	c.JSON(200, gin.H{
		"code":   200,
		"status": "success",
		"data": gin.H{
			"product":  product,
			"ratings":  ratings,
			"recomend": recomend,
		},
	})
}

func AvgRating(ratings []database.Review) float64 {
	avg := 0.0
	if len(ratings) == 0 {
		return avg
	}
	for _, v := range ratings {
		avg += v.Rating
	}
	return math.Round((avg/float64(len(ratings)))*100) / 100
}
