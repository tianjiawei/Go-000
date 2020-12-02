package dao

import (
	"Go-000/Week02/db/mysql"
	err "github.com/pkg/errors"
)

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
	if result.Error != nil || result.RowsAffected <= 0 {
		return user, err.Wrap(result.Error, "user dao query error!")
	}
	return user, nil
}
