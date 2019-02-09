package servers

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func initMiddlewares(e *echo.Echo) {
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
}
