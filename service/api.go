package service

import (
	"github.com/Camelia-hu/im/db"
	"github.com/Camelia-hu/im/model"
	"github.com/gin-gonic/gin"
)

// @Summary 测试swagger
// @Produce json
// Param name query string false "用户名称"
// @Success 200 {object} model.UserBasic "请求成功"
// @Router /service/api.go [post]
func Ping(c *gin.Context) {
	str := c.Query("name")
	var user model.UserBasic
	db.DB.Where("name = ?", str).First(&user)
	c.JSON(200, user)
}
