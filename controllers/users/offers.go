package users

import (
	"ecom/database"
	"ecom/helper"
)

func ProductOffer(id interface{}) float64 {
	var offer database.Offers
	if err := helper.DB.Preload("Product").Where("product_id=?", id).First(&offer); err.Error != nil {
		return 0
	}
	resutl := (offer.Product.ProductPrice / 100) * offer.Percentage
	return resutl
}
