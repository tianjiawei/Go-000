package dao

import (
	"Go-000/Week04/internal/db/mysql"
	"Go-000/Week04/internal/model"
)

type UserDao struct {
	db *mysql.UserDB
}

func NewUserDao(db *mysql.UserDB) *UserDao {
	return &UserDao{
		db: db,
	}
}

func (u *UserDao) GetUserById(uid int64) (user model.User) {
	u.db.Db.Where("id = ?", uid).Find(&user)
	return
}
