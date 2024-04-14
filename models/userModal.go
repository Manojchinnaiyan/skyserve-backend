package models

import "gorm.io/gorm"

type User struct{
	gorm.Model
	Email string `gorm:"unique"`
	Password string
    GeoData []GeoJSON `gorm:"foreignKey:UserID"`
}

type GeoJSON struct {
    gorm.Model
    UserID      uint 
	GeoJsonData string
}