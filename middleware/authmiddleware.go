package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Manojchinnaiyan/initalizers"
	"github.com/Manojchinnaiyan/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)


func AuthMiddleware(c *gin.Context){
	tokenString, err := c.Cookie("token")
	
	if(err != nil){
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(os.Getenv("JWT_SECRET")), nil
   })
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	fmt.Println(token)

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if(float64(time.Now().Unix()) > claims["exp"].(float64)){
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		var user models.User
		initalizers.DB.First(&user, claims["sub"])
		if(user.ID == 0){
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Set("user",user)
		c.Next()
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}