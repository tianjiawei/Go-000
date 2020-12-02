package dao

import "Go-000/Week02/db/mysql"

type User struct {
	Uid     int64  `json:"uid"`
	Name    string `json:"name"`
	Avatar  string `json:"avatar"`
	Created int    `json:"created"`
}

type UserDao struct {
	UserDB *mysql.UserDB
}

func NewUserDao(userDb *mysql.UserDB) *UserDao {
	return &UserDao{
		UserDB: userDb,
	}
}

func (u *UserDao) GetUserById(uid int64) (User, error) {
	var user User
	result := u.UserDB.Db.Table("wx_user_test").Where("uid = ?", uid).First(&user)
	if result.Error == nil {
		return user, nil
	}
	return user, result.Error
}
