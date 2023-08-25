package routes

import (
	"medpicBack/controller"
	"medpicBack/repository"
	"medpicBack/services"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func Router(app *fiber.App, db *mongo.Database) {
	medpicRepo := repository.NewMedpicRepository(db)
	medpicSrv := services.NewMedpicServices(medpicRepo)
	medpicRest := controller.NewMedpicRest(medpicSrv)

	app.Post("/api/v1/medpic_create", medpicRest.Create)
	app.Get("/api/v1/medpic_query", medpicRest.GetQuery)
}
