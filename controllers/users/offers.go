package users

import (
	"ecom/database"
	"ecom/helper"
)

func ProductOffer(id interface{}) float64 {
	var offer database.Product
	if err := helper.DB.Preload("Offers").Where("id=?", id).First(&offer); err.Error != nil {
		return 0
	}
	result := (offer.ProductPrice / 100) * offer.Offers.Percentage
	return result
}
