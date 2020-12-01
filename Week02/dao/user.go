package dao

import "Go-000/Week02/db/mysql"

type User struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Age     string `json:"age"`
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

func (u *UserDao) GetUserById(uid int64) (*User, error) {
	var user *User
	result := u.UserDB.Db.Table("wx_user").Where("uid = ?", uid).First(&user)
	if result.Error == nil {
		return user, nil
	}
	return user, result.Error
}
