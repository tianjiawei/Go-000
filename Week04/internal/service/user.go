package service

import (
	"Go-000/Week04/internal/dao"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"strconv"
)

type UserSrv struct {
	userDao *dao.UserDao
}

func NewUserSrv(userDao *dao.UserDao) *UserSrv {
	return &UserSrv{
		userDao: userDao,
	}
}

func (us *UserSrv) GetUserInfo(c *gin.Context) {
	uid := c.Request.FormValue("uid")
	if uid == "" {
		c.JSON(http.StatusBadRequest, "uid is empty")
	}
	uidInt, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, "uid is wrong!")
	}
	user := us.userDao.GetUserById(uidInt)
	user.PassWord = ""

	userJson, _ := jsoniter.Marshal(user)
	c.JSON(http.StatusOK, userJson)
	return
}
