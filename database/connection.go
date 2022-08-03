package database

import (
	"fmt"

	"github.com/D-Toshchakov/React-JWT-tutorial/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"os"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()

	dbUser := os.Getenv("DBUSER")
	dbPassword := os.Getenv("DBPASSWORD")
	dbName := os.Getenv("DBNAME")
	connString := fmt.Sprintf("%s:%s@/%s",dbUser, dbPassword, dbName)
	connection, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
