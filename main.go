package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"azure-labs-test/app/routers"
	"azure-labs-test/config"
	"azure-labs-test/database"
)

// @title AzureLab Test
// @version 2.0
// @description This is a testing server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	config.InitDB()
	db := config.DB
	database.InitMigration(db)
	database.InitSeeder((db))

	routers.Init(app)
	app.Listen(":" + os.Getenv("APP_PORT"))

}
