package admin

import (
	"ecom/database"
	"ecom/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Sammary Admin can add coupons
// @Description  Admin can add the coupon in ecommerce website
// @Tags Admin-Coupons
// @Accept       multipart/form-data
// @Produce      json
// @Param code formData string true "Coupon Code"
// @Param amount formData string true "Coupon Amount"
// @Param limit formData string true "Limit Amount"
// @Success      200  {string} string "Coupon added"
// @Router       /admin/coupon [post]
func AddCoupons(c *gin.Context) {

	amount, _ := strconv.ParseFloat(c.Request.FormValue("amount"), 64)
	limit, _ := strconv.ParseFloat(c.Request.FormValue("limit"), 64)
	coupon := database.Coupon{
		Code:   c.Request.FormValue("code"),
		Amount: amount,
		Limit:  uint(limit),
	}

	helper.DB.Create(&coupon)
	c.JSON(200, gin.H{"Message": "Coupen added.", "Status": 200})
}

// @Sammary Admin can get all coupons
// @Description Admin side list all the coupons
// @Tags Admin-Coupons
// @Produce      json
// @Success      200  {string} string "Coupons"
// @Router       /admin/coupon [get]
func Coupons(c *gin.Context) {
	var coupons []database.Coupon

	helper.DB.Find(&coupons)
	var list []gin.H
	for _, v := range coupons {
		list = append(list, gin.H{
			"Id":           v.ID,
			"Limit Amount": v.Limit,
			"Coupon Code":  v.Code,
			"Amount":       v.Amount,
		})
	}
	c.JSON(200, gin.H{"Coupons": list, "Status": 200})
}

// @Summary Admin can delete a coupon
// @Description Admin side delete a coupon
// @Tags Admin-Coupons
// @Produce json
// @Param ID path string true "Coupon ID"
// @Success 200 {string} string "Coupon deleted"
// @Router /admin/coupon/{ID} [delete]
func DeleteCoupon(c *gin.Context) {
	Id := c.Param("ID")

	helper.DB.Where("id=?", Id).Delete(&database.Coupon{})
	c.JSON(200, gin.H{"Message": "Coupon deleted."})
}
