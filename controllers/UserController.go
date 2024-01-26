package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Zaw-Thet-Paing/API_V1/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func (u UserController) GetUsers(c echo.Context) error {

	var users []models.User

	result := u.DB.Find(&users)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, users)
}

func (u UserController) GetUser(c echo.Context) error {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var user models.User

	result := u.DB.First(&user, id)

	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"status": "Not found user by this id",
		})
	}

	return c.JSON(http.StatusOK, user)
}

func (u UserController) CreateUser(c echo.Context) error {

	var user models.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad request"})
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	result := u.DB.Create(&user)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusCreated, map[string]string{"status": "created"})
}

func (u UserController) UpdateUser(c echo.Context) error {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var user models.User
	var existing_user models.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad request"})
	}

	user.UpdatedAt = time.Now()

	result := u.DB.First(&existing_user, id)

	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"status": "User not found by this id",
		})
	}

	existing_user.Name = user.Name
	existing_user.Email = user.Email
	existing_user.Password = user.Password
	existing_user.UpdatedAt = user.UpdatedAt

	result = u.DB.Save(&existing_user)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Internal server error",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "updated success"})
}

func (u UserController) DeleteUser(c echo.Context) error {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var existing_user models.User
	result := u.DB.First(&existing_user, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"status": "Not found user by this id",
		})
	}

	u.DB.Delete(&existing_user)

	return c.JSON(http.StatusOK, map[string]string{
		"status": "user deleted success",
	})
}
