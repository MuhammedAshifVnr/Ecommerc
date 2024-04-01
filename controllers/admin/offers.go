package admin

import (
	"ecom/database"
	"ecom/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

func OfferAdding(c *gin.Context) {
	var offer, form database.Offers
	id, _ := strconv.Atoi(c.Param("ID"))
	c.ShouldBindJSON(&form)
	helper.DB.Where("product_id=?", id).First(&offer)
	if offer.ID == 0 {
		form.ProductID = uint(id)
		helper.DB.Create(&form)
		c.JSON(200, gin.H{"Massage": "offer added."})
	} else {
		helper.DB.Model(&offer).Updates(&form)
		c.JSON(200, gin.H{"Massage": "offer updated."})
	}

}

func OfferListing(c *gin.Context){
	var offers []database.Offers

	helper.DB.Find(&offers)
	c.JSON(200,gin.H{"Offers":offers})
}

func OfferDeleting(c *gin.Context){
	id:=c.Param("ID")
	var offer database.Offers
	helper.DB.Where("product_id=?",id).Delete(&offer)
	c.JSON(200,gin.H{"Massage":"Offer deleted."})
}