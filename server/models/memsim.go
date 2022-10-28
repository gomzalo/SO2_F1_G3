package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Memsim struct {
	UserId   primitive.ObjectID `json:"user_id"`
	Ciclos   int   `json:"ciclos"`
	Unidades []int `json:"unidades"`
}