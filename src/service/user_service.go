package service

import (
	"github.com/jinzhu/gorm"
	"github.com/nnnk7777/go-gin-gorm-mysql/model"
)

type UserService struct{}

type User model.User

// GetUserList ユーザ一覧を取得
func (s UserService) GetUserList(db *gorm.DB) ([]User, error) {
	var u []User
	if err := db.Find(&u).Error; err != nil {
		return nil, err
	}
	return u, nil

}
