package controller

import (
	"Go-000/Week02/service"
	"github.com/gin-gonic/gin"
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
	if e != nil {
		//todo
	}
	data, e := u.UserSrv.GetUserById(uid)
	resData := make(map[string]interface{})
	if e != nil {
		resData["code"] = 400
		resData["msg"] = "return is empty！"
	} else {
		resData["code"] = 200
		resData["msg"] = "success！"
	}
	resData["data"] = data
	c.JSON(http.StatusOK, resData)
}
