package users

import (
	"context"
	"ecom/database"
	"ecom/helper"
	"ecom/middleware"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	oauthStateString = "random"
)

// @Summary Login
// @Description Login
// @Tags User
// @Accept  json
// @Produce  json
// @Param user body database.Logging true "Login"
// @Router /user/login [post]
func Login(c *gin.Context) {
	var user database.Logging
	var Find database.User
	c.ShouldBindJSON(&user)
	helper.DB.First(&Find, "email=?", user.Username)

	err := bcrypt.CompareHashAndPassword([]byte(Find.Password), []byte(user.Password))
	if err != nil {
		c.JSON(400, gin.H{"code": 401, "status": "error", "message": "Invalid Username or Password."})
		return
	}
	if Find.Status == "Block" {
		c.JSON(400, gin.H{"code": 401, "status": "error", "message": "User Blocked"})
		return
	}

	token, err := middleware.GenerateToken(Find.Role, Find.Email, Find.ID, Find.Name)
	if err != nil {
		fmt.Println("TOken cant generate.")
	}
	c.SetCookie("user", token, int((time.Hour * 24).Seconds()), "/", "ashif.online", false, false)
	fmt.Println(token)
	c.JSON(200, gin.H{
		"code":    200,
		"status":  "success",
		"message": "successfully Login.",
		"data": gin.H{
			"token": token,
		},
	})

}

// @Summary Logout
// @Description Logout
// @Tags User
// @Produce json
// @Success 200 {object} string "ok"
// @Router /user/logout [get]
func Logout(c *gin.Context) {
	c.SetCookie("user", "", -1, "/", "ashif.online", false, false)
	c.JSON(200, gin.H{"code": 200, "status": "Success", "message": "Successfully logout.", "data": gin.H{}})

}

func GetConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  "http://ashif.online/user/auth/google/callback",
		Scopes:       []string{"openid", "profile", "email"},
		Endpoint:     google.Endpoint,
	}
}

func HandleGoogleLogin(c *gin.Context) {
	googleOauthConfig := GetConfig()
	fmt.Println("+++++++")
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func HandleGoogleCallback(c *gin.Context) {
	googleOauthConfig := GetConfig()
	state := c.Request.URL.Query().Get("state")
	if state != oauthStateString {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid state"})
		return
	}

	fmt.Println("===========0======")
	code := c.Request.URL.Query().Get("code")
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("===========1======")
	client := googleOauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("===========2======")
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("usre---", userInfo)
	c.JSON(http.StatusOK, userInfo)
}
