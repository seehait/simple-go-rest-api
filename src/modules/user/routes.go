package user

import (
	"github.com/labstack/echo"
)

// InitRoutes bind each route with associated handler.
func InitRoutes(e *echo.Echo) {
	e.POST("/user/create", RegisterUser)
	e.POST("/user/login", Login)
}
