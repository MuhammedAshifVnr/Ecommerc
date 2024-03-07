package users

import (
	"ecom/database"
	"ecom/helper"
	"strings"

	"github.com/gin-gonic/gin"
)

func SeaechProduct(c *gin.Context) {
	searchQuery := c.Query("query")
	sortBy := strings.ToLower(c.DefaultQuery("sort", "a_to_z"))

	query := helper.DB
	if searchQuery != "" {
		query = query.Where("product_name ILIKE ?", "%"+searchQuery+"%")
	}

	switch sortBy {
	case "price_low_to_high":
		query = query.Order("product_price asc")
	case "price_high_to_low":
		query = query.Order("product_price desc")
	case "new_arrivals":
		query = query.Order("created_at desc")
	case "a_to_z":
		query = query.Order("product_name asc")
	case "z_to_a":
		query = query.Order("product_name desc")
	case "popularity":
		var products []database.Product
		query := `SELECT * FROM products
				JOIN (
					SELECT
						product_id,
						SUM(quantity) as total_quantity
					FROM
						orders
					GROUP BY
						product_id
					ORDER BY
						total_quantity DESC
					LIMIT 10
				) AS subq ON products.id = subq.product_id
				WHERE
					products.deleted_at IS NULL
				ORDER BY
					subq.total_quantity DESC`
		helper.DB.Raw(query).Scan(&products)

		for _, v := range products {
			c.JSON(200, gin.H{
				"Name":     v.ProductName,
				"Price":    v.ProductPrice,
				"Category": v.Category.Name,
				"ID":       v.ID,
			})
		}
		return
	default:
		query = query.Order("product_name asc")
	}
	var items []database.Product
	query.Find(&items)

	for _, v := range items {
		c.JSON(200, gin.H{
			"Name":     v.ProductName,
			"Price":    v.ProductPrice,
			"Category": v.Category.Name,
			"ID":       v.ID,
		})
	}
}
