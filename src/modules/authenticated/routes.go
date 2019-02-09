package authenticated

import (
	"simple-go-rest-api/configs"
	"simple-go-rest-api/src/modules/user"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// InitRoutes bind each route with associated handler.
func InitRoutes(e *echo.Echo) {
	g := e.Group("/authenticated")
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: configs.GetJwtConfig().Secret,
		Claims:     &user.JwtStruct{},
	}))

	g.GET("", Authenticated)
}
