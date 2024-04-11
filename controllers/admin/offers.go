package admin

import (
	"ecom/database"
	"ecom/helper"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Offer Adding
// @Description Product Offer Adding
// @Tags Admin-Offer
// @Accept  json
// @Produce  json
// @Param ID path int true "Product ID"
// @Param offer body database.OfferProductData true "Offer"
// @Router /admin/offer/{ID} [post]
func OfferAdding(c *gin.Context) {
	var offer database.Offers
	var form database.OfferProductData
	id, _ := strconv.Atoi(c.Param("ID"))
	c.ShouldBindJSON(&form)
	helper.DB.Where("product_id=?", id).First(&offer)
	fmt.Println(form)
	if offer.ID == 0 {
		helper.DB.Create(&database.Offers{
			ProductID:  uint(id),
			Percentage: form.Percentage,
			Expirey:    form.Expirey,
		})
		c.JSON(200, gin.H{"code": 200, "status": "Success", "message": "offer added."})
	} else {
		helper.DB.Model(&offer).Updates(&form)
		c.JSON(200, gin.H{"code": 200, "status": "Success", "message": "offer updated."})
	}

}

// @Summary Product Offer Listing
// @Description Product Offer Listing
// @Tags Admin-Offer
// @Produce  json
// @Router /admin/offer [get]
func OfferListing(c *gin.Context) {
	var offers []database.Offers

	helper.DB.Find(&offers)
	var offer_list []gin.H
	for _, offer := range offers {
		offer_list = append(offer_list, gin.H{
			"ID":         offer.ID,
			"ProductID":  offer.ProductID,
			"Percentage": offer.Percentage,
			"Expirey":    offer.Expirey,
		})

	}
	c.JSON(200, gin.H{"code": 200, "status": "Success", "Offers": offer_list})
}

// @Summary Product Offer Deleting
// @Description Product Offer Deleting
// @Tags Admin-Offer
// @Produce  json
// @Param ID path int true "Product ID"
// @Router /admin/offer/{ID} [delete]
func OfferDeleting(c *gin.Context) {
	id := c.Param("ID")
	var offer database.Offers
	helper.DB.Where("product_id=?", id).Delete(&offer)
	c.JSON(200, gin.H{"code": 200, "status": "Success", "message": "Offer deleted."})
}
