package main

import (
	conf "Go-000/Week02/config"
	"Go-000/Week02/controller"
	"Go-000/Week02/dao"
	"Go-000/Week02/db/mysql"
	handler "Go-000/Week02/http"
	"Go-000/Week02/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	//httpHandle
	container.Provide(handler.NewHttp)

	//controller
	container.Provide(controller.NewHandler)
	container.Provide(controller.NewUserController)
	//config
	container.Provide(conf.NewConfig)
	//db
	container.Provide(mysql.NewUserDB)
	//dao
	container.Provide(dao.NewUserDao)
	//service
	container.Provide(service.NewUserSrv)
	return container
}

var Db *gorm.DB

func main() {
	container := BuildContainer()

	//start http server
	err := container.Invoke(func(router *gin.Engine) {
		router.Run(":8080")
	})
	if err != nil {
		panic(err)
	}
}
