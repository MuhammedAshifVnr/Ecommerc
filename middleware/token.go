package middleware

import (
	"ecom/database"
	"ecom/helper"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var Jwtkey = []byte("secret_Key")

type Claims struct {
	UserID   uint
	UserMail string
	Role     string
	UserName string
	jwt.StandardClaims
}

func GenerateToken(role string, email string, id uint, name string) (string, error) {
	claims := Claims{
		id,
		email,
		role,
		name,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(Jwtkey)
}

func AuthMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString ,err:= c.Cookie(role)
		if err!=nil{
			c.JSON(401,gin.H{"Massage":"Token not found go to the login"})
			c.Abort()
			return
		}
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"erorr": "Unauthorized"})
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) { return Jwtkey, nil })
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"erorr": "Invalid or expired token."})
			c.Abort()
			return
		}
		claims, ok := token.Claims.(*Claims)
		if !ok {
			c.JSON(500, gin.H{"error": "Failed to calling claims."})
		}
		if claims.Role != role {
			c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission."})
			c.Abort()
			return
		}
		var user database.User
		helper.DB.Where("id=?", claims.UserID).First(&user)
		if user.Status == "Block" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Blocked User."})
			c.Abort()
			return
		}
		c.Set("userID", claims.UserID)
		c.Set("useremail", claims.UserMail)
		c.Set("role", claims.Role)
		c.Set("username", claims.UserName)
		c.Next()
	}
}
