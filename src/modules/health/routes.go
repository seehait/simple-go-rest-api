package health

import (
	"github.com/labstack/echo"
)

// InitRoutes bind each route with associated handler.
func InitRoutes(e *echo.Echo) {
	e.GET("/health", Health)
}
