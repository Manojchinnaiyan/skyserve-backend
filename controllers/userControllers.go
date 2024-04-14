package controllers

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Manojchinnaiyan/initalizers"
	"github.com/Manojchinnaiyan/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)


func SignUp(c *gin.Context){
	var body struct{
		Email string 
		Password string
		ConfirmPassword string
	}

	if c.Bind(&body) != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Fail to read nbody",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Failed to has the password",
		})
		return
	}

	user := models.User{Email: body.Email, Password: string(hash)}
	result:= initalizers.DB.Create(&user)

	if result.Error != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Failed to create user",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	"sub": user.ID,
	"exp": time.Now().Add(time.Hour * 24 *30).Unix(),
     })

   tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
   if err != nil{
	c.JSON(http.StatusBadRequest, gin.H{
		"error":"Invalis to create token",
	})
   }

	c.JSON(http.StatusOK, gin.H{
        "token":tokenString,
	})
}

func SignIn(c *gin.Context){
		var body struct{
		Email string
		Password string
	}

	if c.Bind(&body) != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Fail to read nbody",
		})
		return
	}

	var user models.User
	initalizers.DB.First(&user, "email = ?", body.Email)
	 if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Invalid username or password",
		})
		return
	 }

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Invalid username or passwrd",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	"sub": user.ID,
	"exp": time.Now().Add(time.Hour * 24 *30).Unix(),
     })

   tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
   if err != nil{
	c.JSON(http.StatusBadRequest, gin.H{
		"error":"Invalis to create token",
	})
	return
   }

   c.SetSameSite(http.SameSiteLaxMode)
   c.SetCookie("Authorization",tokenString, 3600*20*30, "", "", false, true)

   c.JSON(http.StatusOK, gin.H{
	"token": tokenString,
	"user": user,
   })
}

func CreateGeoData(c *gin.Context){
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error":"Something went wrong",
		})
		return
	}
    geodata := models.GeoJSON{UserID: 10,GeoJsonData: string(jsonData)}
	initalizers.DB.Create(&models.GeoJSON{})

	c.JSON(http.StatusOK, gin.H{
		"ok":geodata,
	})
}


func GetGeodata(c *gin.Context){
	var geodata []models.GeoJSON
	initalizers.DB.Find(&geodata)

	c.JSON(http.StatusOK, gin.H{
		"geodata": geodata,
	})

}
