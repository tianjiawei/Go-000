package service

import (
	"Go-000/Week02/dao"
	"github.com/pkg/errors"
)

type UserSrv struct {
	UserDao *dao.UserDao
}

func NewUserSrv(userDao *dao.UserDao) *UserSrv {
	return &UserSrv{
		UserDao: userDao,
	}
}

func (u *UserSrv) GetUserById(id int64) (map[string]interface{}, error) {
	user, e := u.UserDao.GetUserById(id)
	data := make(map[string]interface{})
	if e != nil {
		return data, errors.Wrap(e, "user service is error!")
	}

	data["uid"] = user.Uid
	data["name"] = user.Name
	data["avatar"] = user.Avatar

	return data, nil
}
