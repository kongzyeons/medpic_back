package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MedpicRepository interface {
	Create(drugs []interface{}) (err error)
	GetQuery(listQuery []bson.M) (results []DrugSimiDB, err error)
	Predict()
}

type medpicRepository struct {
	db *mongo.Database
}

func NewMedpicRepository(db *mongo.Database) MedpicRepository {
	return medpicRepository{db}
}

func (obj medpicRepository) Create(drugs []interface{}) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = obj.db.Collection("drug_similarity").InsertMany(ctx, drugs)
	if err != nil {
		return err
	}
	return err
}

func (obj medpicRepository) GetQuery(listQuery []bson.M) (results []DrugSimiDB, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := obj.db.Collection("drug_similarity").Aggregate(ctx, listQuery)
	if err != nil {
		return results, err
	}
	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &results); err != nil {
		return results, err
	}
	return results, err

}

func (obj medpicRepository) GetAll() (drugs DrugSimiDB, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := obj.db.Collection("drug_similarity").Find(ctx, bson.M{})
	if err = cursor.All(context.TODO(), &drugs); err != nil {
		return drugs, err
	}
	return drugs, err
}

func (obj medpicRepository) Predict() {}
