package users

import (
	"ecom/database"
	"ecom/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Adding Product in cart
// @Description User can add product in cart
// @Tags       User-Cart
// @Produce    json
// @Param ID path int true "Product ID"
// @Router /user/cart/{ID} [post]
func AddCart(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("ID"))
	var cart database.Cart
	var product database.Product
	userId := c.GetUint("userID")
	helper.DB.Where("product_id=? AND user_id=?", id, userId).First(&cart)
	if cart.ID != 0 {
		c.JSON(400, gin.H{"code": 400, "status": "error", "data": gin.H{}, "message": "Item alredy in cart."})
		return
	}
	helper.DB.Where("id=?", id).First(&product)
	if product.Quantity == 0 {
		c.JSON(400, gin.H{"code": 400, "status": "error", "data": gin.H{}, "message": "Product out of stock"})
		return
	}
	cart = database.Cart{
		UserID:    userId,
		ProductID: uint(id),
		Quantity:  1,
	}
	helper.DB.Create(&cart)
	c.JSON(200, gin.H{"code": 200, "status": "Success", "data": gin.H{}, "message": "added to cart."})
}

// @Summary    Get cart
// @Description User can get the cart
// @Tags       User-Cart
// @Produce    json
// @Router /user/cart [get]
func Cart(c *gin.Context) {
	var cart []database.Cart
	helper.DB.Preload("Product.Offers").Where("user_id=?", c.GetUint("userID")).Find(&cart)
	var total float64
	var discount float64
	if len(cart) == 0 {
		c.JSON(401, gin.H{"Massege": "Cart is empty"})
		return
	}
	var cart_list []gin.H
	for _, v := range cart {
		discount += ProductOffer(v.ProductID)
		total += (float64(v.Product.ProductPrice) - ProductOffer(v.ProductID)) * float64(v.Quantity)
		cart_list = append(cart_list, gin.H{
			"ID":           v.ID,
			"productName":  v.Product.ProductName,
			"offerPrice":   v.Product.ProductPrice - ProductOffer(v.ProductID),
			"quantity":     v.Quantity,
			"productImage": v.Product.ImageUrls,
			"orginalPrice": v.Product.ProductPrice,
			"offer":        v.Product.Offers.Percentage,
		})
	}
	var delvery int
	if total <= 1500 {
		delvery = 40
		total += float64(delvery)
	}
	c.JSON(200, gin.H{"code": 200, "status": "Success", "data": gin.H{
		"Cart":          cart_list,
		"Discount":      discount,
		"DliveryCharge": delvery,
		"TotalAmount":   total,
	}})
}

// @Summary Delete in Cart
// @Description User can delete the product from cart
// @Tags User-Cart
// @Accept json
// @Produce json
// @Param ID path string true "ID"
// @Router /user/cart/{ID} [delete]
func CartDelete(c *gin.Context) {
	id := c.Param("ID")
	var cart database.Cart
	helper.DB.Where("id=?", id).Delete(&cart)
	c.JSON(200, gin.H{"code": 200, "status": "Success", "data": gin.H{}, "message": "Deleted."})
}

// @Summary Update Quantity in Cart
// @Description User can update the quantity of product in cart
// @Tags User-Cart
// @Accept multipart/form-Data
// @Produce json
// @Param ID path string true "ID"
// @Param quantity formData int true "Quantity"
// @Router /user/cart/{ID} [patch]
func CartQuantity(c *gin.Context) {
	id := c.Param("ID")
	var cart database.Cart
	quantity, _ := strconv.Atoi(c.Request.FormValue("quantity"))

	helper.DB.Preload("Product").Where("id=?", id).First(&cart)

	if quantity > 10 {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "status": "error", "data": gin.H{}, "message": "Only 10 unit's allowed each order"})
		return
	}
	if quantity > cart.Product.Quantity {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "status": "error", "data": gin.H{}, "message": "Product is out of stock"})
		return
	}
	if quantity <= 0 {
		c.JSON(400, gin.H{"code": 400, "status": "error", "data": gin.H{}, "message": "Please Enter a valid Quantity."})
		return
	}

	helper.DB.Model(&database.Cart{}).Where("id=?", id).Update("Quantity", quantity)
	c.JSON(200, gin.H{"code": 200, "status": "Success", "data": gin.H{}, "message": "Quantity Updated."})
}
