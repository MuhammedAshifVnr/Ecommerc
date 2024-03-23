package users

import (
	"ecom/database"
	"ecom/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Whislist(c *gin.Context) {
	var list []database.Whislist
	helper.DB.Preload("Product").Where("user_id=?", c.GetUint("userID")).Find(&list)
	if len(list)==0{
		c.JSON(502,gin.H{
			"massage:":"your whislist is empty",
		})
		return
	}
	c.JSON(200, gin.H{
		"Products": list,
	})
}

func AddWhislist(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("ID"))
	list := database.Whislist{
		UserID:    c.GetUint("userID"),
		ProductID: uint(id),
	}
	helper.DB.Create(&list)
	c.JSON(200, gin.H{
		"message": "Product added",
	})
}

func DeleteWhislist(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("ID"))
	helper.DB.Where("id=?", id).Delete(&database.Whislist{})
	c.JSON(200, gin.H{
		"massage": "Product deleted.",
	})
}
