//+build wireinject

package main

import (
	"Go-000/Week04/internal/dao"
	"Go-000/Week04/internal/db/mysql"
	"Go-000/Week04/internal/service"
	"github.com/google/wire"
)

func InitUserSrv() service.UserSrv {
	wire.Build(service.NewUserSrv, dao.NewUserDao)
	return service.UserSrv{}
}

func InitUserDao() dao.UserDao {
	wire.Build(dao.NewUserDao, mysql.NewUserDB)
	return dao.UserDao{}
}
