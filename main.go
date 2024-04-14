package main

import (
	"github.com/Manojchinnaiyan/controllers"
	"github.com/Manojchinnaiyan/initalizers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init(){
	initalizers.LoadEnvVariables()
	initalizers.ConnectToDB()
	initalizers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/signup",controllers.SignUp )
	r.POST("/signin", controllers.SignIn)
	r.POST("/createGeo", controllers.CreateGeoData)
	r.GET("/getGeodata", controllers.GetGeodata)
	r.Run()
}