package admin

import (
	"ecom/database"
	"ecom/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UsersList(c *gin.Context) {
	var Table []database.User
	helper.DB.Order("ID").Find(&Table)

	for _, v := range Table {
		c.JSON(http.StatusSeeOther, gin.H{
			"ID":     v.ID,
			"Name":   v.Name,
			"Email":  v.Email,
			"Status": v.Status,
		})
	}

}

func UserStatus(c *gin.Context) {
	var find database.User
	id := c.Param("ID")
	helper.DB.First(&find, id)

	if find.Status == "Active" {
		find.Status = "Block"
		helper.DB.Save(&find)
		c.JSON(http.StatusAccepted, "User Blocked")
	} else {
		find.Status = "Active"
		helper.DB.Save(&find)
		c.JSON(http.StatusAccepted, "User Unblocked")
	}

}
