package users

import (
	"ecom/database"
	"ecom/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddCart(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("ID"))
	var cart database.Cart
	helper.DB.Where("product_id=? AND user_id=?", id, Find.ID).First(&cart)
	if cart.ID != 0 {
		c.JSON(400, "Item alredy in cart.")
		return
	}
	cart = database.Cart{
		UserID:    Find.ID,
		ProductID: uint(id),
		Quantity:  1,
	}
	helper.DB.Create(&cart)
	c.JSON(200, "added to cart.")
}

func Cart(c *gin.Context) {
	var cart []database.Cart
	helper.DB.Preload("Product").Where("user_id=?", Find.ID).Find(&cart)
	for _, v := range cart {

		c.JSON(200, v)
	}
	var total float64
	for _, v := range cart {
		total += float64(v.Product.ProductPrize) * float64(v.Quantity)
	}
	c.JSON(200, gin.H{
		"Total Amount": total,
	})
}

func CartDelete(c *gin.Context) {
	id := c.Param("ID")
	var cart database.Cart
	helper.DB.Where("id=?", id).Delete(&cart)
	c.JSON(200, "Deleted.")
}

func CartQuantity(c *gin.Context) {
	id := c.Param("ID")
	var cart database.Cart
	quantity, _ := strconv.Atoi(c.Request.FormValue("quantity"))

	helper.DB.Preload("Product").Where("id=?", id).First(&cart)
	if quantity > cart.Product.Quantity {
		c.JSON(http.StatusBadRequest, "Product is out of stock")
		return
	}
	if quantity > 10 {
		c.JSON(http.StatusBadRequest, "Only 10 unit's allowed each order")
		return
	}

	helper.DB.Model(&database.Cart{}).Where("id=?", id).Update("Quantity", quantity)
	c.JSON(200, "Quantity Updated.")
}
