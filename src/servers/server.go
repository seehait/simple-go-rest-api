package servers

import "github.com/labstack/echo"

// Init server.
func Init() *echo.Echo {
	e := echo.New()

	initMiddlewares(e)
	initRoutes(e)

	return e
}
