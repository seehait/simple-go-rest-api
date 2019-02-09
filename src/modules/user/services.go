package user

import (
	"fmt"
	"net/http"
	"simple-go-rest-api/configs"
	"simple-go-rest-api/src/databases"
	"simple-go-rest-api/src/databases/models"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func encryptPassword(password string) ([]byte, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError)
	}

	return encryptedPassword, nil
}

func createUser(name string, email string, password string) (*models.User, error) {
	encryptedPassword, err := encryptPassword(password)
	if err != nil {
		return nil, err
	}

	user := models.User{Name: name, Email: email, Password: string(encryptedPassword)}
	database := databases.Manager()
	database.Create(&user)

	return &user, nil
}

func emailIsNotBeUsed(email string) bool {
	var existingUser models.User

	database := databases.Manager()
	if database.Where("email = ?", email).First(&existingUser).RecordNotFound() {
		return true
	}

	return false
}

func validateLogin(email string, password string) (*models.User, error) {
	var user models.User
	database := databases.Manager()
	if database.Where("email = ?", email).First(&user).RecordNotFound() {
		return nil, echo.NewHTTPError(http.StatusUnauthorized)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, echo.NewHTTPError(http.StatusUnauthorized)
	}

	return &user, nil
}

// JwtStruct represents jwt structure.
type JwtStruct struct {
	jwt.StandardClaims
}

func generateJwtToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtStruct{jwt.StandardClaims{
		Id:        fmt.Sprintf("%d", userID),
		ExpiresAt: time.Now().Unix() + 15000,
	}})

	tokenString, err := token.SignedString(configs.GetJwtConfig().Secret)
	if err != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError)
	}

	return tokenString, nil
}

// GetUserIDFromContext extracts user id from context.
func GetUserIDFromContext(c echo.Context) (uint64, error) {
	token := c.Get("user")
	if token, ok := token.(*jwt.Token); ok == true {

		claims := token.Claims.(*JwtStruct)
		userID, err := strconv.ParseUint(claims.StandardClaims.Id, 10, 64)
		if err != nil {
			return 0, echo.NewHTTPError(http.StatusInternalServerError)
		}

		return userID, nil
	}

	return 0, echo.NewHTTPError(http.StatusUnauthorized)
}
