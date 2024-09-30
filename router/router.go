package router

import (
	_ "github.com/Camelia-hu/im/docs"
	"github.com/Camelia-hu/im/service"
	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()
	r.POST("/ping", service.Ping)

	r.Run()
}
