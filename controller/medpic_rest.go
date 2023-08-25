package controller

import (
	"log"
	"medpicBack/models"
	"medpicBack/services"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type MedpicRest interface {
	Create(c *fiber.Ctx) error
	GetQuery(c *fiber.Ctx) error
}

type medpicRest struct {
	medpicSrv services.MedpicServices
}

func NewMedpicRest(medpicSrv services.MedpicServices) MedpicRest {
	return medpicRest{medpicSrv}
}

func (obj medpicRest) Create(c *fiber.Ctx) error {
	model := models.CreateRequest{}
	if err := c.BodyParser(&model); err != nil {
		log.Println("error", err)
		return c.Status(http.StatusBadRequest).JSON(models.Err_response(err))
	}
	if err := validator.New().Struct(&model); err != nil {
		log.Println("error", err)
		return c.Status(http.StatusBadRequest).JSON(models.Err_response(err))
	}
	err := obj.medpicSrv.Create(model.Drugs)
	if err != nil {
		log.Println("error", err)
		return c.Status(http.StatusInternalServerError).JSON(models.Err_response(err))
	}
	log.Println("success", c.Status(http.StatusCreated))
	var data []interface{}
	return c.Status(http.StatusCreated).JSON(models.Response(data, len(data), "success"))
}

func (obj medpicRest) GetQuery(c *fiber.Ctx) error {
	model := models.GetQueryRequest{}
	if err := c.BodyParser(&model); err != nil {
		log.Println("error", err)
		return c.Status(http.StatusBadRequest).JSON(models.Err_response(err))
	}
	if err := validator.New().Struct(&model); err != nil {
		log.Println("error", err)
		return c.Status(http.StatusBadRequest).JSON(models.Err_response(err))
	}
	data, err := obj.medpicSrv.GetQuery(model.Query)
	if err != nil {
		log.Println("error", err)
		return c.Status(http.StatusInternalServerError).JSON(models.Err_response(err))
	}
	log.Println("success", c.Status(http.StatusOK))
	return c.Status(http.StatusOK).JSON(models.Response(data, len(data), "success"))
}
