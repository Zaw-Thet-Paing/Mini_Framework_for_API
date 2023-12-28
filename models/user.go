package models

import (
	"github.com/Zaw-Thet-Paing/API_V1/models/entities"
	"gorm.io/gorm"
)

type UserModel struct {
	Db *gorm.DB
}

func (userModel UserModel) FindAll() ([]entities.User, error) {
	db := userModel.Db

	var users []entities.User

	result := db.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil

}

func (userModel UserModel) FindById(id int64) entities.User {
	db := userModel.Db

	var user entities.User

	db.Where("id = ?", id).First(&user)

	return user
}
