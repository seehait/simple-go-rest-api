package servers

import (
	"simple-go-rest-api/src/modules/authenticated"
	"simple-go-rest-api/src/modules/health"
	"simple-go-rest-api/src/modules/user"

	"github.com/labstack/echo"
)

func initRoutes(e *echo.Echo) {
	authenticated.InitRoutes(e)
	health.InitRoutes(e)
	user.InitRoutes(e)
}
