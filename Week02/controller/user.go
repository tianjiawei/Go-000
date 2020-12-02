package controller

import (
	"Go-000/Week02/log"
	"Go-000/Week02/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type UserController struct {
	UserSrv *service.UserSrv
}

func NewUserController(userSrv *service.UserSrv) *UserController {
	return &UserController{
		UserSrv: userSrv,
	}
}

func (u *UserController) GetUser(c *gin.Context) {
	uidStr := c.Request.FormValue("uid")
	uid, e := strconv.ParseInt(uidStr, 10, 64)
	resData := make(map[string]interface{})
	if e != nil {
		resData["code"] = 400
		resData["msg"] = "uid is empty！"
		c.JSON(http.StatusOK, resData)
	}
	data, err := u.UserSrv.GetUserById(uid)
	if err != nil {
		fmt.Printf("stack trace:\n%+v\n", err)
		log.Logger.Error("error", zap.String("stack trace:", fmt.Sprintf("\n%+v\n", err)))
		resData["code"] = 400
		resData["msg"] = "info is empty！"
	} else {
		resData["code"] = 200
		resData["msg"] = "success！"
	}
	resData["data"] = data
	c.JSON(http.StatusOK, resData)
}
