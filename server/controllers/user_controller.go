package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"server/configs"
	"server/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "usuarios")

func CreateUser(w http.ResponseWriter, r *http.Request) {	
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	
	var user models.User
	var responseMessage models.Message

	w.Header().Set("Content-Type", "application/json")
	
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		responseMessage.Message = "Por favor, ingrese datos con la cuenta de usuario"
		json.NewEncoder(w).Encode(responseMessage)
		return
	}
	
	json.Unmarshal(reqBody, &user)
	newUser := models.User {
		Id: primitive.NewObjectID(),
		Email: user.Email,
		Pass: user.Pass,
	}
	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		responseMessage.Message = "Error al crear cuenta"
		json.NewEncoder(w).Encode(responseMessage)
		return
	}
	fmt.Printf("%v", result.InsertedID)
	w.WriteHeader(http.StatusCreated)
	responseMessage.Message = "Cuenta creada satisfactoriamente"
	responseMessage.IdUser = newUser.Id
	json.NewEncoder(w).Encode(responseMessage)
}

func GetOneUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	
	var responseMessage models.Message
	var findedUser models.User
	var searchUser models.User

	w.Header().Set("Content-Type", "application/json")

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		responseMessage.Message = "Por favor, ingrese datos con la cuenta de usuario"
		json.NewEncoder(w).Encode(responseMessage)
		return
	}

	json.Unmarshal(reqBody, &searchUser)
	
	err = userCollection.FindOne(ctx, bson.M{"email": searchUser.Email, "pass": searchUser.Pass}).Decode(&findedUser)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		responseMessage.Message = "Usuario no encontrado"
		json.NewEncoder(w).Encode(responseMessage)
		return
	}
	
	w.WriteHeader(http.StatusOK)
	responseMessage.Message = "Acceso aprobado"
	json.NewEncoder(w).Encode(findedUser)
}


func GetAllUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	var users []models.User
	var responseMessage models.Message
	defer cancel()

	results, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		responseMessage.Message = "No se encontraron usuarios"
		json.NewEncoder(w).Encode(responseMessage)
		return
	}
	
	defer results.Close(ctx)
	for results.Next(ctx) {
		var user models.User
		if err = results.Decode(&user); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			responseMessage.Message = "Error del servidor"
			json.NewEncoder(w).Encode(responseMessage)
			return
		}
		users = append(users, user)
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
