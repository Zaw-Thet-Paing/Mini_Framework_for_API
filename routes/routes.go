package routes

import (
	"github.com/Zaw-Thet-Paing/API_V1/controllers"
	"github.com/labstack/echo/v4"
)

func Setup(e *echo.Echo) {

	//user controller
	user := controllers.UserController{}
	e.GET("/", user.GetUsers)
}
