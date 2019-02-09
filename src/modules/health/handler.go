package health

import (
	"net/http"

	"github.com/labstack/echo"
)

type healthResponse struct {
	Status string `json:"status"`
}

// Health handle health checking.
// It always return "OK".
func Health(c echo.Context) error {
	return c.JSON(http.StatusOK, &healthResponse{Status: "OK"})
}
