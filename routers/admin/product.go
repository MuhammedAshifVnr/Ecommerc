package admin

import (
	"ecom/database"
	"ecom/helper"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type prodectbind struct {
	database.Product
	Categ string `json:"catagory"`
}

var prod database.Product
var edited database.Product

// ..................................product showing................................
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

// .............................admin can add products......................
func AddProduct(c *gin.Context) {
	var find prodectbind
	var cate database.Category

	c.ShouldBindJSON(&find)

	helper.DB.Where("name=?", find.Categ).First(&cate)

	prod = database.Product{
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
	//helper.DB.Create(&prod)
	c.JSON(200, "Please Upload the Product Images.")

}

// .....................................product image adding...............................................
func ImageAdding(c *gin.Context) {

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
		return
	}

	file := form.File["image"]

	if len(file) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error2": "Minimum 3 images required"})
		return
	}

	for i := 0; i < len(file); i++ {

		dst := filepath.Join("./assets", file[i].Filename)
		err := c.SaveUploadedFile(file[i], dst)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error3": err.Error()})
			return
		}
		if i == 0 {
			prod.ImageUrl1 = dst
		} else if i == 1 {
			prod.ImageUrl2 = dst
		} else {
			prod.ImageUrl3 = dst
		}
	}
	helper.DB.Create(&prod)
	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}

// ......................................admin can edit the product..............................
func EditProdect(c *gin.Context) {
	var bind prodectbind
	var category database.Category

	c.ShouldBind(&bind)
	id := c.Param("ID")
	helper.DB.Where("name=?", bind.Categ).First(&category)
	edited = database.Product{
		ProductName:  bind.ProductName,
		CategoryId:   category.ID,
		Quantity:     bind.Quantity,
		Size:         bind.Size,
		ProductPrize: bind.ProductPrize,
		Description:  bind.Description,
	}
	if edited.CategoryId == 0 {
		c.JSON(200, "Category not found Please give a valid Category.")
		return
	}
	helper.DB.Model(&database.Product{}).Where("id=?", id).Updates(edited)
	c.JSON(200, "Redirecting to Image edit.")
}

// ......................product image editing................................
func ImageEditing(c *gin.Context) {
	edited=database.Product{}
	id := c.Param("ID")
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
		return
	}

	file := form.File["image"]

	if len(file) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error2": "Minimum 3 images required"})
		return
	}

	for i := 0; i < len(file); i++ {

		dst := filepath.Join("./assets", file[i].Filename)
		err := c.SaveUploadedFile(file[i], dst)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error3": err.Error()})
			return
		}
		if i == 0 {
			prod.ImageUrl1 = dst
		} else if i == 1 {
			prod.ImageUrl2 = dst
		} else {
			prod.ImageUrl3 = dst
		}
	}
	helper.DB.Model(&database.Product{}).Where("id=?", id).Updates(edited)
	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
	c.JSON(200, "Successfully Edited.")
}

// .................................admin can delete product............................
func Delete(c *gin.Context) {
	var delete database.Product

	id := c.Param("ID")

	helper.DB.Where("id=?", id).First(&delete)
	if err := helper.DB.Where("id=?", id).Delete(&delete); err.Error != nil {
		c.JSON(http.StatusSeeOther, "You Can't Delete this Product.")
	}
	c.JSON(http.StatusSeeOther, "Successfylly Deleted.")

}
