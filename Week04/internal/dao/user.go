package dao

import (
	"Go-000/Week04/internal/model"
	"github.com/jinzhu/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{
		db: db,
	}
}

func (u *UserDao) GetUserById(uid int64) (user model.User) {
	return
}
