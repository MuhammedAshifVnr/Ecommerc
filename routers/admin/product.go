package admin

import (
	"ecom/database"
	"ecom/helper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type prodectbind struct {
	database.Product
	Categ string `json:"catagory"`
}

func Product(c *gin.Context) {
	var product []database.Product
	helper.DB.Order("ID").Find(&product)

	for _, v := range product {
		c.JSON(http.StatusSeeOther, gin.H{
			"Product Name": v.ProductName,
			"Category":     v.CategoryId,
			"Quantity":     v.Quantity,
			"Prize":        v.ProductPrize,
			"Status":       v.Status,
		})
	}

}

func AddProduct(c *gin.Context) {
	var find prodectbind
	var cate database.Category

	c.ShouldBindJSON(&find)

	helper.DB.Where("name=?", find.Categ).First(&cate)
	fmt.Println(cate)
	fmt.Println(
		find.ProductName,
		"---", find.Categ,
		find.Quantity,
		find.Size,
		find.ProductPrize,
		find.Description,
	)

	prod := database.Product{
		ProductName:  find.ProductName,
		CategoryId:   cate.ID,
		Quantity:     find.Quantity,
		Size:         find.Size,
		ProductPrize: find.ProductPrize,
		Description:  find.Description,
	}
	if prod.CategoryId == 0 {
		c.JSON(200, "Category not found Please give a valid Category.")
		return
	}
	helper.DB.Create(&prod)
	c.JSON(200, "Product added Successfully.")

}

func EditProdect(c *gin.Context) {

}

func Delete(c *gin.Context) {
	var delete database.Product

	id := c.Param("ID")

	helper.DB.Where("id=?", id).First(&delete)
	if err := helper.DB.Where("id=?", id).Delete(&delete); err.Error != nil {
		c.JSON(http.StatusSeeOther, "Can't Delete this user.")
	}
	c.JSON(http.StatusSeeOther, "Successfylly Delete User.")

}
