package service

import (
	"Go-000/Week04/internal/dao"
	"github.com/gin-gonic/gin"
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
		return
	}
	uidInt, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, "uid is wrong!")
		return
	}
	user := us.userDao.GetUserById(uidInt)
	userObj := make(map[string]interface{})
	userObj["id"] = user.Id
	userObj["user_name"] = user.UserName
	userObj["age"] = user.Age
	c.JSON(http.StatusOK, userObj)
	return
}
