package server

import (
	"Go-000/Week04/internal/service"
	"github.com/gin-gonic/gin"
)

func NewHttp() *gin.Engine {
	router := gin.Default()

	userSrv := service.InitUserSrv()
	router.GET("/get/user", userSrv.GetUserInfo)
	return router
}
