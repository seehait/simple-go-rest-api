package databases

import (
	"fmt"

	"simple-go-rest-api/configs"
	"simple-go-rest-api/src/databases/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var database *gorm.DB
var err error

// Init initializes database manager.
func Init() {
	databaseConfiguration := configs.GetDatabaseConfig()
	connectString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", databaseConfiguration.UserName, databaseConfiguration.Password, databaseConfiguration.Host, databaseConfiguration.Port, databaseConfiguration.Name)
	database, err = gorm.Open("mysql", connectString)

	if err != nil {
		panic("Database Connection Error")
	}

	database.AutoMigrate(&models.User{})
}

// Manager provides database manager.
func Manager() *gorm.DB {
	return database
}
