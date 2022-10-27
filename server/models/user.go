package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id    primitive.ObjectID `json:"id,omitempty"`
	Email string `json:"email"`
	Pass  string `json:"pass"`
}