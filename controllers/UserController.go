package controllers

import (
	"log"
	"net/http"

	"github.com/Zaw-Thet-Paing/API_V1/databases"
	"github.com/Zaw-Thet-Paing/API_V1/models"
	"github.com/labstack/echo/v4"
)

type UserController struct {
}

func (u UserController) GetUsers(c echo.Context) error {

	db, err := databases.Connect()
	if err != nil {
		log.Fatal(err)
	}

	userModel := models.UserModel{Db: db}
	users, err := userModel.FindAll()

	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, users)
}
