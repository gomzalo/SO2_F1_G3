package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	Message  string `json:"message"`
	IdUser primitive.ObjectID `json:"id,omitempty"`
}