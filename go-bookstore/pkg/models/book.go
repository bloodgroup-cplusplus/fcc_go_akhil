package models 

import (
	"github.com/jinzhu/gorm"
	"github.com/bloodgroupcplusplus/fcc_go_akhil/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.model
	Name string `gorm:""json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}


func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})

}

// control is first in the routes and the routes give control to controllers and controllers will give controls to the models 
// we'll have to have different functions for differnt functions
// different functions 



