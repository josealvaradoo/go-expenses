package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/josealvaradoo/expenses/database"
	"github.com/josealvaradoo/expenses/models"
	"github.com/josealvaradoo/expenses/router"
	"github.com/josealvaradoo/expenses/utils"
)

func initDatabaseConnection() {
	var err error

	environment := utils.AccessEnv("APP_ENV")
	user := utils.AccessEnv("DB_USER")
	dbname := utils.AccessEnv("DB_NAME")
	password := utils.AccessEnv("DB_PASS")
	port := utils.AccessEnv("DB_PORT")
	host := "localhost"

	if environment == "production" {
		host = "postgres"
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	database.DBConnection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// Migrations
	database.DBConnection.AutoMigrate(&models.User{})
	database.DBConnection.AutoMigrate(&models.Currency{})

	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("Database connection was successfully")
}

func main() {
	// Init database

	// Init app
	app := fiber.New()

	// Init database
	initDatabaseConnection()

	// Config routes
	router.SetupRoutes(app)

	// Serve
	app.Listen(8000)
}
