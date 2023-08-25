package services

import (
	"fmt"
	"medpicBack/models"
	"medpicBack/repository"
	"medpicBack/utils"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
)

type MedpicServices interface {
	Create(drugs []repository.DrugSimiDB) (err error)
	GetQuery(listStr []utils.FieldValueBson) (results []models.GetQueryResponse, err error)
}

type medpicServices struct {
	medpicRepo repository.MedpicRepository
}

func NewMedpicServices(medpicRepo repository.MedpicRepository) MedpicServices {
	return medpicServices{medpicRepo}
}

func (obj medpicServices) Create(drugs []repository.DrugSimiDB) (err error) {
	drugsInterface := []interface{}{}
	for i := range drugs {
		if err := validator.New().Struct(drugs[i]); err != nil {
			return err
		}
		listQuery := []bson.M{
			{
				"$match": bson.M{
					"name": drugs[i].Name,
				},
			},
			{
				"$match": bson.M{
					"label": drugs[i].Label,
				},
			},
			{
				"$match": bson.M{
					"color": drugs[i].Color,
				},
			},
		}
		dbDrugs, err := obj.medpicRepo.GetQuery(listQuery)
		if err != nil {
			return err
		}
		if len(dbDrugs) != 0 {
			err = fmt.Errorf("name drug not ready")
			return err
		}

		drugs[i].CreateDate = time.Now()
		drugs[i].UpdateDate = time.Now()
		drugsInterface = append(drugsInterface, drugs[i])
	}

	err = obj.medpicRepo.Create(drugsInterface)
	if err != nil {
		return err
	}
	return err
}

func (obj medpicServices) GetQuery(listStr []utils.FieldValueBson) (results []models.GetQueryResponse, err error) {
	if len(listStr) != 0 {
		for i := range listStr {
			fmt.Println(i)
			if err := validator.New().Struct(listStr[i]); err != nil {
				return results, err
			}
		}
	}
	listQuery := utils.ListStr2Bson(listStr)
	drugs, err := obj.medpicRepo.GetQuery(listQuery)
	if err != nil {
		return results, err
	}
	for i := range drugs {
		result := models.GetQueryResponse{
			Name:      drugs[i].Name,
			Color:     drugs[i].Color,
			Label:     drugs[i].Label,
			Embedding: drugs[i].Embedding,
		}
		results = append(results, result)
	}

	return results, err
}
