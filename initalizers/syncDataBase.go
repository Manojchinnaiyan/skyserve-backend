package initalizers

import "github.com/Manojchinnaiyan/models"

func SyncDatabase(){
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.GeoJSON{})
}