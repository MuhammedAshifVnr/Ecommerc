package admin

import (
	"ecom/database"
	"ecom/helper"

	"github.com/gin-gonic/gin"
)

// @Summary      get all category list by admin
// @Description  Admin can see the listed  category in ecommerce website
// @Tags Admin-Category
// @Produce      json
// @Router       /admin/category [get]
func Category(c *gin.Context) {
	var find []database.Category

	helper.DB.Find(&find)
	var category []gin.H
	for _, v := range find {
		category = append(category, gin.H{
			"ID":          v.ID,
			"Name":        v.Name,
			"Description": v.Description,
			"Status":      v.Status,
		})
	}
	c.JSON(200, gin.H{
		"Category": category,
		"Status":   200,
	})
}

// @Summary Add a new category
// @Description Add a new category with the provided details
// @ID add-category
// @Tags Admin-Category
// @Accept json
// @Produce json
// @Param category body database.CategoryData true "Category Add"
// @Success 200 {string} string "Successfully added."
// @Failure 404 {string} string "Duplicate found."
// @Router       /admin/category [post]
func AddCategory(c *gin.Context) {
	var cat database.CategoryData
	c.ShouldBindJSON(&cat)

	if err := helper.DB.Create(&database.Category{
		Name:        cat.Name,
		Description: cat.Description,
	}); err.Error != nil {
		c.JSON(404, gin.H{"Message": "Duplicate found."})
		return
	}
	c.JSON(200, gin.H{"Message": "Successfully added."})

}

// @Summary      Admin can edit the category
// @Description  Admin can edit the listed  category in ecommerce website
// @Tags Admin-Category
// @Produce      json
// @Param ID path string true "User ID"
// @Param Form body database.CategoryData true "Category Edit"
// @Router       /admin/category/{ID} [put]
func EditCategory(c *gin.Context) {
	var find database.Category

	id := c.Param("ID")

	c.ShouldBindJSON(&find)
	if err := helper.DB.Where("id=?", id).Updates(&find); err.Error != nil {
		c.JSON(404, "error")
		return
	}
	c.JSON(200, gin.H{"Message": "Successfully Edited."})
}

// @Summary      Admin can block the category
// @Description  Admin can Block the listed  category in ecommerce website
// @Tags Admin-Category
// @Produce      json
// @Param ID path string true "User ID"
// @Router       /admin/category/{ID} [patch]
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
		c.JSON(401, gin.H{"Message": "Successfully Blocked.", "Status": 401})
	} else {
		helper.DB.Model(&find).Update("Status", "Active")
		helper.DB.Where("category_id=?", id).Find(&prod)
		for _, v := range prod {
			helper.DB.Model(&v).Update("status", "Active")
		}
		c.JSON(200, gin.H{"Message": "Successfully Actived.", "Status": 200})
	}
}

// @Summary Delete a category
// @Description Deletes a category by ID
// @Tags Admin-Category
// @Produce json
// @Param ID path string true "Category ID"
// @Router /admin/category/{ID} [delete]
func DeleteCategory(c *gin.Context) {
	var find database.Category
	var prod database.Product

	id := c.Param("ID")

	helper.DB.Where("CategoryId=?", id).First(&prod)
	if prod.ID != 0 {
		c.JSON(422, gin.H{"Message": "You can't Delete this Category.Some Product Listed this Category.", "Status": 422})
		return
	}

	helper.DB.Where("id=?", id).Delete(&find)
	c.JSON(200, gin.H{"Message": "Category deleted Successfully.", "Status": 200})
}

func DeleteRecovery(c *gin.Context) {
	id := c.Param("ID")

	helper.DB.Unscoped().Model(&database.Category{}).Where("id=?", id).Update("deleted_at", nil)
	c.JSON(200, "Recoverd.")
}
