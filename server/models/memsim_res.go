package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type MemsimRes struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	UserId   primitive.ObjectID `json:"user_id,omitempty"`
	Memsim   []MemsimProc `json:"memsim"`
	Duracion int64        `json:"duracion"`
	Procesos int          `json:"procesos"`
	Unidades []int          `json:"unidades"`
}