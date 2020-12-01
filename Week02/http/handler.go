package http

import (
	"Go-000/Week02/controller"

	"github.com/gin-gonic/gin"
)

func NewHttp(handle *controller.Handler) *gin.Engine {
	router := gin.Default()

	router.POST("/user/get", handle.UserController.GetUser)
	return router
}
