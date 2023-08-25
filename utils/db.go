package utils

import "go.mongodb.org/mongo-driver/bson"

type FieldValueBson struct {
	Field string `json:"field" validate:"required"`
	Value string `json:"value" validate:"required"`
	Key   string `json:"key" validate:"required"`
}

func ListStr2Bson(listQuery []FieldValueBson) (retsults []bson.M) {
	for i := range listQuery {
		retsult := bson.M{}
		retsult[listQuery[i].Key] = bson.M{
			listQuery[i].Field: listQuery[i].Value,
		}
		retsults = append(retsults, retsult)
	}
	return retsults
}
