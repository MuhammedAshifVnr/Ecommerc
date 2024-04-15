package users

import (
	"ecom/database"
	"ecom/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary User can write review
// @Description User can write review
// @Tags User-Review
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param review body database.ReviewData true "Review"
// @Router /user/review/{id} [post]
func CreatReview(c *gin.Context) {
	var input database.ReviewData
	//var order database.OrderItems
	strId, _ := c.Get("userID")
	userID, _ := strId.(uint)
	id, _ := strconv.ParseUint(c.Param("ID"), 10, 64)
	// if err := helper.DB.Where("product_id=? AND user_id=?", id, userID).First(&order); err.Error != nil {
	// 	c.JSON(400, gin.H{"code": 400, "status": "Failed", "message": "We cant add the review.", "data": gin.H{}})
	// 	return
	// }
	c.ShouldBindJSON(&input)

	helper.DB.Create(&database.Review{
		Rating:    input.Rating,
		Comment:   input.Comment,
		UserID:    userID,
		ProductID: uint(id),
	})
	c.JSON(200, gin.H{"code": 200, "status": "Success", "message": "review added.", "data": gin.H{}})
}
