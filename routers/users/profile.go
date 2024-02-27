package users

import (
	"ecom/database"
	"ecom/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {
	c.JSON(200, gin.H{
		"Name":          Find.Name,
		"Email":         Find.Email,
		"Mobile Number": Find.Mobile,
	})
}

func AddAddress(c *gin.Context) {
	var address database.Address

	c.ShouldBindJSON(&address)
	address.UserId = Find.ID
	helper.DB.Create(&address)
	c.JSON(http.StatusCreated, "Address added.")

}

func Address(c *gin.Context){
	var add []database.Address

	helper.DB.Where("user_id=?",Find.ID).Find(&add)

	for _, v := range add {
		c.JSON(200,v)
	}
}

func AddressEdit (c *gin.Context){
	id := c.Param("ID")
	var add database.Address
	c.ShouldBindJSON(&add)
	add.UserId=Find.ID
	helper.DB.Model(&database.Address{}).Where("id=?",id).Updates(add)
	c.JSON(200,"Successfully Edited.")
}

func AddressDelete (c *gin.Context){
	id := c.Param("ID")
	var add database.Address
	helper.DB.Where("id=?",id).Delete(&add)
}