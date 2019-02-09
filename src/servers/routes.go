package servers

import (
	"simple-go-rest-api/src/modules/health"

	"github.com/labstack/echo"
)

func initRoutes(e *echo.Echo) {
	health.InitRoutes(e)
}
