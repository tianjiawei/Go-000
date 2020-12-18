//+build wireinject

package main

import (
	"Go-000/Week04/configs"
	"Go-000/Week04/internal/dao"
	"Go-000/Week04/internal/db/mysql"
	"Go-000/Week04/internal/service"
	"github.com/google/wire"
)

func InitUserSrv() *service.UserSrv {
	panic(wire.Build(service.NewUserSrv, dao.NewUserDao, mysql.NewUserDB, configs.NewConfig))
}
