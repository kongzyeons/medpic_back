package repository

import "time"

type DrugSimiDB struct {
	DrugID     string    `json:"drug_id" bson:"drug_id"`
	Embedding  []float64 `josn:"embedding" bson:"embedding" validate:"required"`
	Color      string    `josn:"color" bson:"color" validate:"required"`
	Label      string    `josn:"label" bson:"label" validate:"required"`
	Name       string    `josn:"name" bson:"name" validate:"required"`
	CreateDate time.Time `json:"create_date" bson:"create_date"`
	UpdateDate time.Time `json:"update_date" bson:"update_date"`
}
