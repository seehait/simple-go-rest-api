package user

import (
	"net/http"

	"github.com/labstack/echo"
)

type userRegistrationRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type userRegistrationResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// RegisterUser handles user registration.
func RegisterUser(c echo.Context) error {
	registrationRequest := new(userRegistrationRequest)
	if err := c.Bind(registrationRequest); err != nil {
		return err
	}

	if !emailIsNotBeUsed(registrationRequest.Email) {
		return echo.NewHTTPError(http.StatusBadRequest, "This email is already used.")
	}

	user, err := createUser(registrationRequest.Name, registrationRequest.Email, registrationRequest.Password)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &userRegistrationResponse{ID: user.ID, Name: user.Name, Email: user.Email})
}
