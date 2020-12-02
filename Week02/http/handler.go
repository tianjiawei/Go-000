package http

import (
	"Go-000/Week02/controller"
	middleware "Go-000/Week02/midware"

	"github.com/gin-gonic/gin"
)

func NewHttp(handle *controller.Handler) *gin.Engine {
	router := gin.Default()

	router.POST("/user/get", middleware.Recover(), middleware.RequestIdHandler(), handle.UserController.GetUser)
	return router
}
