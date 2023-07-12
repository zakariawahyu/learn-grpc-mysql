package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func ConnectDatabase() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:masukdb@tcp(localhost:3306)/go_grpc_mysql"))
	if err != nil {
		log.Fatal("Can't connect database %v", err.Error())
	}

	return db
}
