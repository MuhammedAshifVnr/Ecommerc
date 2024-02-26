package admin

import (
	"ecom/database"
	"ecom/helper"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type prodectbind struct {
	database.Product
	Categ string `json:"catagory"`
}

// var prod database.Product
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
			"images":v.ImageUrls,
		})
	}

}


func AddProduct(c *gin.Context) {

	var cate database.Category

	helper.DB.First(&cate).Where("name=?", c.Request.FormValue("category"))

	quantit, _ := strconv.Atoi(c.Request.FormValue("quantity"))
	prize, _ := strconv.ParseFloat(c.Request.FormValue("prize"), 64)
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
		ProductPrize: prize,
		CategoryId:   cate.ID,
		Quantity:     quantit,
		Size:         size,
		Description:  c.Request.FormValue("description"),
		ImageUrls:    pq.StringArray(imgs),
	}
	fmt.Println("-------------------------------", imgs, "===============")
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
// func ImageEditing(c *gin.Context) {
// 	edited = database.Product{}
// 	id := c.Param("ID")
// 	form, err := c.MultipartForm()
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
// 		return
// 	}

// 	file := form.File["image"]

// 	if len(file) < 3 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error2": "Minimum 3 images required"})
// 		return
// 	}

// 	for i := 0; i < len(file); i++ {

// 		dst := filepath.Join("./assets", file[i].Filename)
// 		err := c.SaveUploadedFile(file[i], dst)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error3": err.Error()})
// 			return
// 		}
// 		if i == 0 {
// 			prod.ImageUrl1 = dst
// 		} else if i == 1 {
// 			prod.ImageUrl2 = dst
// 		} else {
// 			prod.ImageUrl3 = dst
// 		}
// 	}
// 	helper.DB.Model(&database.Product{}).Where("id=?", id).Updates(edited)
// 	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
// 	c.JSON(200, "Successfully Edited.")
// }

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
