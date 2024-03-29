package users

import (
	"ecom/database"
	"ecom/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatReview(c *gin.Context) {
	var input database.Review
	var order database.Order
	strId, _ := c.Get("userID")
	userID, _ := strId.(uint)
	id, _ := strconv.ParseUint(c.Param("ID"), 10, 64)
	if err := helper.DB.Where("product_id=? AND user_id=?", id, userID).First(&order); err.Error != nil {
		c.JSON(400, "We cant add the review.")
	}
	c.ShouldBindJSON(&input)

	input.UserID = userID
	input.ProductID = uint(id)
	helper.DB.Create(&input)
	c.JSON(200, "review added.")
}
