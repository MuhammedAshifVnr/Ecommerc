package users

import (
	"ecom/database"
	"ecom/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {
	var user database.User
	helper.DB.Where("id=?",Find.ID).First(&user)
	c.JSON(200, gin.H{
		"Name":          user.Name,
		"Email":         user.Email,
		"Mobile Number": user.Mobile,
		"Gender":        user.Gender,
	})
}

func AddAddress(c *gin.Context) {
	var address database.Address

	c.ShouldBindJSON(&address)
	address.UserId = Find.ID
	helper.DB.Create(&address)
	c.JSON(http.StatusCreated, "Address added.")

}

func Address(c *gin.Context) {
	var add []database.Address

	helper.DB.Where("user_id=?", Find.ID).Find(&add)

	for _, v := range add {
		c.JSON(200, gin.H{
			"City":v.City,
			"ID":v.ID,
			"State":v.State,
			"Street":v.Street,
			"Type":v.Type,
			"Zip":v.ZipCode,
		})
	}
}

func AddressEdit(c *gin.Context) {
	id := c.Param("ID")
	var add database.Address
	c.ShouldBindJSON(&add)
	add.UserId = Find.ID
	helper.DB.Model(&database.Address{}).Where("id=?", id).Updates(add)
	c.JSON(200, "Successfully Edited.")
}

func AddressDelete(c *gin.Context) {
	id := c.Param("ID")
	var add database.Address
	helper.DB.Where("id=?", id).Delete(&add)
}

func EditeProfile(c *gin.Context) {
	var edite database.User
	helper.DB.Where("id=?", Find.ID).First(&edite)
	c.ShouldBindJSON(&edite)
	if err:=helper.DB.Save(&edite);err.Error !=nil{
		c.JSON(400,"can't save")
	}
	c.JSON(200, "Updated.")
}
