package users

import (
	"ecom/database"
	"ecom/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary User Profile
// @Description Get user profile
// @Tags User-Profile
// @Produce  json
// @Router /user/profile [get]
func Profile(c *gin.Context) {
	var user database.User
	helper.DB.Where("id=?", c.GetUint("userID")).First(&user)
	profile := gin.H{
		"Name":          user.Name,
		"Email":         user.Email,
		"Mobile Number": user.Mobile,
		"Gender":        user.Gender,
	}
	c.JSON(200, gin.H{
		"code":   200,
		"status": "success",
		"data":   profile,
	})
}

// @Summary Address Adding
// @Description User can Add Address
// @Tags User-Address
// @Accept json
// @Produce  json
// @Param address body database.AddressData true "Address"
// @Router /user/address [post]
func AddAddress(c *gin.Context) {
	var address database.Address
	userId := c.GetUint("userID")
	c.ShouldBindJSON(&address)
	address.UserId = userId
	helper.DB.Create(&address)
	c.JSON(http.StatusCreated, gin.H{"code": 201, "status": "Success", "message": "Address added.", "data": gin.H{}})
}

// @Summary Address listing
// @Description User can list Address
// @Tags User-Address
// @Accept json
// @Produce  json
// @Router /user/address [get]
func Address(c *gin.Context) {
	var add []database.Address
	userId := c.GetUint("userID")
	helper.DB.Where("user_id=?", userId).Find(&add)
	var address_list []gin.H
	for _, v := range add {
		address_list = append(address_list, gin.H{
			"City":   v.City,
			"ID":     v.ID,
			"State":  v.State,
			"Street": v.Street,
			"Type":   v.Type,
			"Zip":    v.ZipCode,
		})
	}
	c.JSON(200, gin.H{
		"code":   200,
		"status": "success",
		"data":   gin.H{"address": address_list},
	})
}

// @Summary Address Edit
// @Description User can Edite Address
// @Tags User-Address
// @Accept json
// @Produce json
// @Param ID path string true "Address ID"
// @Param user body database.AddressData true "User"
// @Router /user/address/{ID} [patch]
func AddressEdit(c *gin.Context) {
	id := c.Param("ID")
	var add database.Address
	userId := c.GetUint("userID")
	c.ShouldBindJSON(&add)
	add.UserId = userId
	helper.DB.Model(&database.Address{}).Where("id=?", id).Updates(add)
	c.JSON(200, gin.H{"code": 200, "status": "success", "message": "Successfully Edited.", "data": gin.H{}})
}

// @Summary Address Delete
// @Description User can Delete Address
// @Tags User-Address
// @Produce json
// @Param ID path string true "Address ID"
// @Router /user/address/{ID} [delete]
func AddressDelete(c *gin.Context) {
	id := c.Param("ID")
	var add database.Address
	helper.DB.Where("id=?", id).Delete(&add)
	c.JSON(200, gin.H{"code": 200, "status": "Success", "message": "Successfully Deleted", "data": gin.H{}})
}

// @Summary Edite Profile
// @Description User can Edite Profile
// @Tags User-Profile
// @Accept json
// @Produce json
// @Param user body database.UserData true "User"
// @Router /user/profile [patch]
func EditeProfile(c *gin.Context) {
	var edite database.User
	userId := c.GetUint("userID")
	helper.DB.Where("id=?", userId).First(&edite)
	c.ShouldBindJSON(&edite)
	if err := helper.DB.Save(&edite); err.Error != nil {
		c.JSON(400, gin.H{"code": 400, "status": "error", "message": "can't save", "data": gin.H{}})
	}
	c.JSON(200, gin.H{"code": 400, "status": "error", "message": "Updated", "data": gin.H{}})
}
