package models

import (
	"medpicBack/repository"
	"medpicBack/utils"
)

type CreateRequest struct {
	Drugs []repository.DrugSimiDB `json:"drugs" validate:"required"`
}

type GetQueryRequest struct {
	Query []utils.FieldValueBson `json:"query" validate:"required"`
}

type GetQueryResponse struct {
	Name      string    `json:"name" bson:"name" validate:"required"`
	Color     string    `json:"color" bson:"color" validate:"required"`
	Label     string    `json:"label" bson:"label" validate:"required"`
	Embedding []float64 `json:"embedding" bson:"embedding" validate:"required"`
}
