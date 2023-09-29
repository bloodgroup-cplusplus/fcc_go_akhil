package config 

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"
)


var (
	db *gorm.DB
)


func Connect(){
	// the database here is mysql database 

	d,err := gorm.Open("mysql","chad:chadleb/simplerest?charset=utf8&parseTime=True&loc=local")
	if err != nil {
		panic(err)
	}
	db = d

}

func GetDb() * gorm.DB{
	return db
}



