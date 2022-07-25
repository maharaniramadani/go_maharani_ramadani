package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var JwtSecret string

func InitDB(){

	connectionString:= "root:@/middleware?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString),&gorm.Config{})
	if err != nil{
		panic(err)
	}
	
}

// type Config struct{
// 	DBNAME string
// 	DBUSER string
// 	DBPASS string
// 	DBHOST string
// 	DBPOrt int

// 	JwtSecret string
// }

// var Conf Config

// func Init(){
// 	Conf = Config{
// 		DBNAME: os.Getenv("DBNAME"),
// 	}

// }
