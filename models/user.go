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

func (userModel UserModel) FindById(id int64) (entities.User, error) {
	db := userModel.Db

	var user entities.User

	db.Where("id = ?", id).First(&user)

	return user, nil
}

func (userModel UserModel) InsertUser(user entities.User) (int, error) {
	db := userModel.Db
	result := db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}

	return int(user.ID), nil
}

func (userModel UserModel) UpdateUser(user entities.User, id int64) (int, error) {
	db := userModel.Db

	existingUser, err := userModel.FindById(id)
	if err != nil {
		return 0, err
	}

	existingUser.Name = user.Name
	existingUser.Email = user.Email
	existingUser.Password = user.Password
	existingUser.UpdatedAt = user.UpdatedAt

	result := db.Save(&existingUser)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(result.RowsAffected), nil
}

func (userModel UserModel) DeleteUser(id int64) (int, error) {
	return 1, nil
}
