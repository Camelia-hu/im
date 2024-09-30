package main

import (
	"github.com/Camelia-hu/im/config"
	"github.com/Camelia-hu/im/db"
	"github.com/Camelia-hu/im/router"
)

func main() {
	config.Config_Init()
	db.DB_init()
	router.Router()
}
