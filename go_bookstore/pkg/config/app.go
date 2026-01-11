package config
import (
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
)
var db *gorm.DB 

func Connect(){
	d ,err := gorm.Open("mysql","Vikram:Vikram_123@tcp(127.0.0.1:3306)/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err !=  nil{
		panic(err) //If DB connection fails:App crashes immediately
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}