package db

import (
	"github.com/Camelia-hu/im/config"
	"github.com/Camelia-hu/im/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func DB_init() {
	db, err := gorm.Open(mysql.Open(config.Conf.GetString("data.mysql.dsn")), &gorm.Config{})
	if err != nil {
		log.Fatalf("db init err   " + err.Error())
	}

	err = db.AutoMigrate(&model.UserBasic{})
	if err != nil {
		log.Fatalf("migrate err   " + err.Error())
	}

	DB = db
}
