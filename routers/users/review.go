package users

import (
	"ecom/database"
	"ecom/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreastReview(c *gin.Context) {
	var input database.Review
	id, _ := strconv.ParseUint(c.Param("ID"), 10, 64)
	c.ShouldBindJSON(&input)
	input.UserID = Find.ID
	input.ProductID = uint(id)
	helper.DB.Create(&input)
}
