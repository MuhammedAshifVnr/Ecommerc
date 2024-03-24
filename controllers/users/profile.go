package users

import (
	"ecom/database"
	"ecom/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {
	var user database.User
	helper.DB.Where("id=?",c.GetUint("userID")).First(&user)
	c.JSON(200, gin.H{
		"Name":          user.Name,
		"Email":         user.Email,
		"Mobile Number": user.Mobile,
		"Gender":        user.Gender,
	})
}

func AddAddress(c *gin.Context) {
	var address database.Address
	userId:=c.GetUint("userID")
	c.ShouldBindJSON(&address)
	address.UserId = userId
	helper.DB.Create(&address)
	c.JSON(http.StatusCreated, "Address added.")

}

func Address(c *gin.Context) {
	var add []database.Address
	userId:=c.GetUint("userID")
	helper.DB.Where("user_id=?", userId).Find(&add)

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
	userId:=c.GetUint("userID")
	c.ShouldBindJSON(&add)
	add.UserId = userId
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
	userId:=c.GetUint("userID")
	helper.DB.Where("id=?", userId).First(&edite)
	c.ShouldBindJSON(&edite)
	if err:=helper.DB.Save(&edite);err.Error !=nil{
		c.JSON(400,"can't save")
	}
	c.JSON(200, "Updated.")
}
