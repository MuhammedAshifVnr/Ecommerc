package admin

import (
	"ecom/database"
	"ecom/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddCoupons(c *gin.Context) {
	
	amount, _ := strconv.ParseFloat(c.Request.FormValue("amount"),64)
	coupon :=database.Coupon{
		Code: c.Request.FormValue("code"),
		Amount: amount,
	}
	
	helper.DB.Create(&coupon)
	c.JSON(200,"Coupen added.")
}

func Coupons (c *gin.Context){
	var coupons []database.Coupon

	helper.DB.Find(&coupons)
	for _, v := range coupons {
		c.JSON(200,gin.H{
			"Coupon Code":v.Code,
			"Amount":v.Amount,
		})
	}
}