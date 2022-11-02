package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"server/configs"
	"server/models"
	"strconv"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var memsimCollection *mongo.Collection = configs.GetCollection(configs.DB, "memsim")
var procesos []string

func CreateMemsim(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	var memsim models.Memsim
	var message models.Message

	w.Header().Set("Content-Type", "application/json")

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		message.Message = "Por favor, ingrese datos de inicialización de comando memsim"
		json.NewEncoder(w).Encode(message)
		return
	}

	json.Unmarshal(reqBody, &memsim)
	var result = execMemsim(memsim.Ciclos, memsim.Unidades)

	result.Id = primitive.NewObjectID()
	result.UserId = memsim.UserId

	memsimCollection.InsertOne(ctx, result)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func GetAllMemsim(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	var listMemsim []models.MemsimRes
	var message models.Message
	var user models.User

	w.Header().Set("Content-Type", "application/json")
	
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		message.Message = "Por favor, ingrese datos con la cuenta de usuario"
		json.NewEncoder(w).Encode(message)
	}

	json.Unmarshal(reqBody, &user)

	results, err := memsimCollection.Find(ctx, bson.D{})
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		message.Message = "No se encontraron resultados de comando memsim"
		json.NewEncoder(w).Encode(message)
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var memsim models.MemsimRes
		if err = results.Decode(&memsim); err != nil {
			fmt.Println("Error de servidor")
			w.WriteHeader(http.StatusInternalServerError)
			message.Message = "Error del servidor"
			json.NewEncoder(w).Encode(message)
			return
		}
		fmt.Printf("Memsim: %v", memsim)
		listMemsim = append(listMemsim, memsim)
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(listMemsim)
}


func execMemsim(ciclos int, unidades []string) models.MemsimRes {
	var result models.MemsimRes
	size := len(unidades)
	now := time.Now()
	var process int = 0

	for i := 1; i <= ciclos; i++ {
		var proceso models.MemsimProc
		var wg sync.WaitGroup // Declarando nuestro wait group
		
		proceso.Ciclo = i
		
		for j := size - 1; j >= 0; j-- {
			wg.Add(1) // Indicamos la cantidad de rutinas a esperar
			value := unidades[j]
			go func() {
				defer wg.Done() // Mensaje region critica
				process = process + 1
				work(process, value, size)
			}()
		}
		wg.Wait()
		proceso.Procesos = procesos
		result.Memsim = append(result.Memsim, proceso)
		procesos = nil
	}

	var duracion = time.Since(now).Milliseconds()
	result.Duracion = duracion
	result.Procesos = process
	result.Unidades = unidades

	return result
}

func work(proceso int, unidad string, tam int) {
	var inicioProceso = "El proceso # " + strconv.Itoa(proceso) + ", empezó a trabajar con la unidad: '" + unidad + "'"
	procesos = append(procesos, inicioProceso)
	
	time.Sleep(time.Duration(tam) * time.Millisecond)
	
	var finProceso = "El proceso # " + strconv.Itoa(proceso) + ", terminó de trabajar con la unidad: '" + unidad + "'"
	procesos = append(procesos, finProceso)
}
