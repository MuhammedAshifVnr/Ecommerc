package admin

import (
	"ecom/database"
	"ecom/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Sammary Admin side user list
// @Description  Admin can see the users in ecommerce website
// @Tags Admin-Users
// @Produce      json
// @Router       /admin/users [get]
func UsersList(c *gin.Context) {
	var Table []database.User
	helper.DB.Order("ID").Find(&Table)
	var users []gin.H
	for _, v := range Table {
		users = append(users, gin.H{
			"ID":     v.ID,
			"Name":   v.Name,
			"Email":  v.Email,
			"Status": v.Status,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"Status": 200,
		"Users":  users,
	})
}

// @Sammary Admin side user Block or Unblock
// @Description  Admin can Block and Unblock the users in ecommerce website
// @Tags Admin-Users
// @Param ID path int true "User ID"
// @Produce      json
// @Router       /admin/users/{ID} [patch]
func UserStatus(c *gin.Context) {
	var find database.User
	id := c.Param("ID")
	helper.DB.First(&find, id)

	if find.Status == "Active" {
		find.Status = "Block"
		helper.DB.Save(&find)
		c.JSON(http.StatusAccepted, gin.H{"Massage": "User Blocked"})
	} else {
		find.Status = "Active"
		helper.DB.Save(&find)
		c.JSON(http.StatusAccepted, gin.H{"Massage": "User Unblocked"})
	}

}
