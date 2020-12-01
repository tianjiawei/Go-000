package service

import "Go-000/Week02/dao"

type UserSrv struct {
	UserDao *dao.UserDao
}

func NewUserSrv(userDao *dao.UserDao) *UserSrv {
	return &UserSrv{
		UserDao: userDao,
	}
}

func (u *UserSrv) GetUserById(id int64) (data interface{}) {
	user, e := u.UserDao.GetUserById(id)

	return
}
