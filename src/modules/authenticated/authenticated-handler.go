package authenticated

import (
	"net/http"

	"github.com/labstack/echo"
)

type authenticatedResponse struct {
	Status string `json:"status"`
}

// Authenticated handle authenticated checking.
// It always return "OK".
func Authenticated(c echo.Context) error {
	return c.JSON(http.StatusOK, &authenticatedResponse{Status: "OK"})
}
