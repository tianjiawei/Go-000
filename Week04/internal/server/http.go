package server

import (
	"github.com/gin-gonic/gin"
)

func NewHttp() *gin.Engine {
	router := gin.Default()

	router.POST("/get/user", nil)
	return router
}
