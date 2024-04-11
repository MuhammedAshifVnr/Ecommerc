package admin

import (
	"ecom/database"
	"ecom/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary      Finding best product and category
// @Description  Admin can find best product and category
// @Tags Admin
// @Produce      json
// @Param        sort   query     string  false  "sort"
// @Router       /admin/sort [get]
func ProductSorting(c *gin.Context) {

	sort := c.Query("sort")
	var items []database.Product
	query := helper.DB
	if sort == "" {
		c.JSON(401, gin.H{"Erorr": "give any value"})
		return
	}
	switch sort {
	case "popular_products":
		query.Raw(`
            SELECT products.*
            FROM products
            JOIN (
                SELECT product_id, SUM(quantity) AS total_quantity
                FROM order_items
                GROUP BY product_id
                ORDER BY total_quantity DESC
            ) AS top_products ON products.id = top_products.product_id 
			ORDER BY top_products.total_quantity DESC
			LIMIT 10
        `).Find(&items)
		var products []gin.H
		for _, v := range items {
			products = append(products, gin.H{
				"id":          v.ID,
				"name":        v.ProductName,
				"description": v.Description,
				"price":       v.ProductPrice,
				"images":      v.ImageUrls,
			})
		}
		c.JSON(http.StatusOK, gin.H{"code": 200, "status": "Success", "top_products": products})
	case "categories":
		var categories []database.Category
		query.Raw(`
		SELECT categories.id, categories.name, SUM(order_items.quantity) AS total_quantity
		FROM categories
		JOIN products ON categories.id = products.category_id
		JOIN order_items ON products.id = order_items.product_id
		GROUP BY categories.id
		ORDER BY total_quantity DESC
		LIMIT 10
        `).Find(&categories)
		var categories_list []gin.H
		for _, v := range categories {
			categories_list = append(categories_list, gin.H{
				"id":   v.ID,
				"name": v.Name,
			})
		}
		c.JSON(http.StatusOK, gin.H{"code": 200, "status": "Success", "top_categories": categories_list})
	default:
		c.JSON(401, gin.H{"code": 401, "status": "Failed", "message": "Wrong Query."})
	}
}
