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
func Product(c *gin.Context) {
	var product []database.Product
	helper.DB.Order("ID").Find(&product)

	for _, v := range product {
		c.JSON(200, gin.H{
			"ID":           v.ID,
			"Product Name": v.ProductName,
			"Category":     v.CategoryId,
			"Quantity":     v.Quantity,
			"Prize":        v.ProductPrice,
			"Status":       v.Status,
			"images":       v.ImageUrls,
			"Stock":        v.Quantity,
		})
	}

}

func AddProduct(c *gin.Context) {
	var cate database.Category

	helper.DB.Where("name=?", c.Request.FormValue("category")).First(&cate)

	quantit, _ := strconv.Atoi(c.Request.FormValue("quantity"))
	price, _ := strconv.ParseFloat(c.Request.FormValue("price"), 64)
	size, _ := strconv.Atoi(c.Request.FormValue("size"))

	if cate.ID == 0 {
		c.JSON(http.StatusBadRequest, "Category not fount.")
		return
	}

	files := c.Request.MultipartForm.File["images"]
	if len(files) < 3 {
		c.JSON(http.StatusBadRequest, "Please upload at least 3 images")
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
	c.JSON(200, "File uploaded successfully")
}

// ......................................admin can edit the product..............................
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
		c.JSON(http.StatusBadRequest, "Category not fount.")
		return
	}

	files := c.Request.MultipartForm.File["images"]
	if len(files) < 3 {
		c.JSON(http.StatusBadRequest, "Please upload at least 3 images")
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
	c.JSON(200, "File uploaded successfully")
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

