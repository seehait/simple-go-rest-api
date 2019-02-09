package user

import (
	"net/http"

	"github.com/labstack/echo"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string `json:"token"`
}

// Login handles user login.
func Login(c echo.Context) error {
	request := new(loginRequest)
	if err := c.Bind(request); err != nil {
		return err
	}

	user, err := validateLogin(request.Email, request.Password)
	if err != nil {
		return err
	}

	tokenString, err := generateJwtToken(user.ID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &loginResponse{Token: tokenString})
}
