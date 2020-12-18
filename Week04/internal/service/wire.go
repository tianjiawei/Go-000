//+build wireinject

package service

import (
	"Go-000/Week04/configs"
	"Go-000/Week04/internal/dao"
	"Go-000/Week04/internal/db/mysql"
	"github.com/google/wire"
)

func InitUserSrv() *UserSrv {
	wire.Build(NewUserSrv, dao.NewUserDao, mysql.NewUserDB, configs.NewConfig)
	//wire.Build(mysql.NewUserDB, configs.NewConfig)
	return &UserSrv{}
}

/*
type Message string
type Greeter struct {
	Message Message
}

func NewMessage() Message {
	return Message("Hi there!")
}

func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

func (g Greeter) Greet() Message {
	return g.Message
}

func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

type Event struct {
	Greeter Greeter // <- adding a Greeter field
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func InitializeEvent() Event {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}
}*/
