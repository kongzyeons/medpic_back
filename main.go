package main

import (
	"medpicBack/database"
	"medpicBack/routes"
	"os"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// server := ":7001"
	server := ":"+ os.Getenv("SERVER_PORT")
	host_mongo := os.Getenv("MONGO_URL")
	db_name := os.Getenv("DB_NAME")

	MongoDB := database.NewMongoDB(host_mongo, db_name)

	app := fiber.New()
	routes.Router(app, MongoDB)
	app.Listen(server)
}
