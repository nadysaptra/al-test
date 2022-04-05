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
