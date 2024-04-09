package admin

import (
	"ecom/database"
	"ecom/helper"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

// ..................................product showing................................
// AdminLogin godoc
// @Summary      Admin can see the listed  products in ecommerce website
// @Description  get all product list by admin
// @Tags Admin-Product
// @Produce      json
// @Router       /admin/product [get]
func Product(c *gin.Context) {
	var product []database.Product
	helper.DB.Order("ID").Find(&product)
	var products []gin.H
	for _, v := range product {
		products = append(products, gin.H{
			"ID":          v.ID,
			"ProductName": v.ProductName,
			"Category":    v.CategoryId,
			"Quantity":    v.Quantity,
			"Prize":       v.ProductPrice,
			"Status":      v.Status,
			"images":      v.ImageUrls,
			"Stock":       v.Quantity,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"Status": http.StatusOK,
		"Data":   products,
	})
}

// @Summary Add a new product
// @Description Add a new product to the database
// @Accept multipart/form-data
// @Produce json
// @Tags Admin-Product
// @Param product formData string true "Name of the product"
// @Param category formData string true "Name of the category the product belongs to"
// @Param quantity formData integer true "Quantity of the product"
// @Param price formData number true "Price of the product"
// @Param size formData integer true "Size of the product"
// @Param description formData string true "Description of the product"
// @Param images formData []file true "Images of the product (upload at least 3 images)"
// @Success 200 {string} string "Product added successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/product [post]
func AddProduct(c *gin.Context) {
	var cate database.Category

	helper.DB.Where("name=?", c.Request.FormValue("category")).First(&cate)

	quantit, _ := strconv.Atoi(c.Request.FormValue("quantity"))
	price, _ := strconv.ParseFloat(c.Request.FormValue("price"), 64)
	size, _ := strconv.Atoi(c.Request.FormValue("size"))

	if cate.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest, "Massage": "Category not fount"})
		return
	}

	files := c.Request.MultipartForm.File["images"]
	if len(files) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest, "Massage": "Please upload at least 3 images"})
		return
	}
	var imgs []string
	for _, img := range files {
		dst := filepath.Join("./assets", img.Filename)
		if err := c.SaveUploadedFile(img, dst); err != nil {
			c.JSON(http.StatusBadRequest, err.Error)
			return
		}
		imgs = append(imgs, dst)
	}
	datas := database.Product{
		ProductName:  c.Request.FormValue("product"),
		ProductPrice: price,
		CategoryId:   cate.ID,
		Quantity:     quantit,
		Size:         size,
		Description:  c.Request.FormValue("description"),
		ImageUrls:    pq.StringArray(imgs),
	}
	if err := helper.DB.Create(&datas); err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
		return
	}
	c.JSON(200, gin.H{"Status": 200, "Massage": "File uploaded successfully"})
}

// ......................................admin can edit the product..............................
// @Summary Edit a product
// @Description Edits a product by its ID, including updating its category, quantity, price, size, description, and images
// @ID edit-product
// @Accept multipart/form-data
// @Tags Admin-Product
// @Produce json
// @Param ID path int true "Product ID"
// @Param product formData string true "Name of the product"
// @Param category formData string true "Product category name"
// @Param quantity formData  string true "Product quantity"
// @Param price formData string true "Product price"
// @Param size formData string true "Product size"
// @Param description formData string true "Product description"
// @Param images formData  []file true "Product images"
// @Router /admin/product/{ID} [put]
func EditProdect(c *gin.Context) {
	id := c.Param("ID")
	var cate database.Category

	var datas database.Product
	helper.DB.Where("id=?", id).First(&datas)

	helper.DB.Where("name=?", c.Request.FormValue("category")).First(&cate)

	quantit, _ := strconv.Atoi(c.Request.FormValue("quantity"))
	price, _ := strconv.ParseFloat(c.Request.FormValue("price"), 64)
	size, _ := strconv.Atoi(c.Request.FormValue("size"))

	if cate.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Status": 400, "Massage": "Category not fount"})
		return
	}

	files := c.Request.MultipartForm.File["images"]
	if len(files) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest, "Massage": "Please upload at least 3 images"})
		return
	}
	var imgs []string
	for _, img := range files {
		dst := filepath.Join("./assets", img.Filename)
		if err := c.SaveUploadedFile(img, dst); err != nil {
			c.JSON(http.StatusBadRequest, err.Error)
			return
		}
		imgs = append(imgs, dst)
	}
	datas = database.Product{
		ProductName:  c.Request.FormValue("product"),
		ProductPrice: price,
		CategoryId:   cate.ID,
		Quantity:     quantit,
		Size:         size,
		Description:  c.Request.FormValue("description"),
		ImageUrls:    pq.StringArray(imgs),
	}
	if err := helper.DB.Model(&database.Product{}).Where("id=?", id).Updates(datas); err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
		return
	}
	c.JSON(200, gin.H{"Status": 200, "Massage": "File uploaded successfully"})
}

// .................................admin can delete product............................
// @Summary Delete a product
// @Description Deletes a product by its ID
// @ID delete-product
// @Tags Admin-Product
// @Accept json
// @Produce json
// @Param ID path int true "Product ID"
// @Router /admin/product/{ID} [delete]
func Delete(c *gin.Context) {
	var delete database.Product

	id := c.Param("ID")

	helper.DB.Where("id=?", id).First(&delete)
	if err := helper.DB.Where("id=?", id).Delete(&delete); err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Massage":"You Can't Delete this Product"})
	}
	c.JSON(http.StatusOK, gin.H{"Status": http.StatusOK, "Massage": "Successfylly Deleted"})

}
