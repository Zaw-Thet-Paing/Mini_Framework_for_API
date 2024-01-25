package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Zaw-Thet-Paing/API_V1/databases"
	"github.com/Zaw-Thet-Paing/API_V1/models"
	"github.com/Zaw-Thet-Paing/API_V1/models/entities"
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

func (u UserController) GetUser(c echo.Context) error {
	db, err := databases.Connect()
	if err != nil {
		log.Fatal(err)
	}

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	userModel := models.UserModel{Db: db}
	user, err := userModel.FindById(id)

	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, user)
}

func (u UserController) CreateUser(c echo.Context) error {

	var user entities.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad request"})
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// log.Println(user)

	db, err := databases.Connect()
	if err != nil {
		log.Fatal(err)
	}

	userModel := models.UserModel{Db: db}
	_, err = userModel.InsertUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"status": "created"})
}

func (u UserController) UpdateUser(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var user entities.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad request"})
	}

	user.UpdatedAt = time.Now()

	// log.Println(user)

	db, err := databases.Connect()
	if err != nil {
		log.Fatal(err)
	}

	userModel := models.UserModel{Db: db}
	_, err = userModel.UpdateUser(user, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
