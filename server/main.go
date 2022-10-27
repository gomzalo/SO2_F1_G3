package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"server/configs"
	"server/controllers"
	"server/models"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

// type message struct {
// 	Message string `json:"message"`
// }

// type account struct {
// 	Id primitive.ObjectID `json:"id,omitempty"`
// 	Email string `json:"email"`
// 	Pass string `json:"pass"`
// }

// type memsim struct {
// 	Ciclos int `json:"ciclos"`
// 	Unidades []int `json:"unidades"`
// }

// type memsim_proc struct {
// 	Ciclo int `json:"ciclo"`
// 	Procesos []string `json:"procesos"`
// }

// type memsim_res struct {
// 	Id primitive.ObjectID `json:"id,omitempty"`
// 	Memsim []memsim_proc `json:"memsim"`
// 	Duracion int64 `json:"duracion"`
// }

// type allAccount []account

// var accounts = allAccount{
// 	{
// 		Email: "usuario@example.com",
// 		Pass: "super-password",
// 	},
// }

var procesos []string

func main() {
	configs.ConnectDB()
	
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/logup", controllers.CreateUser)
	router.HandleFunc("/login", controllers.GetOneUser)
	router.HandleFunc("/users", controllers.GetAllUser)
	router.HandleFunc("/memsim", callMemsim)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var responseMessage models.Message
	responseMessage.Message ="Welcome home!"
	json.NewEncoder(w).Encode(responseMessage)
}

// func createUser(w http.ResponseWriter, r *http.Request) {
// 	var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "usuarios")
	
// 	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
// 	var user account
// 	defer cancel()

// 	w.Header().Set("Content-Type", "application/json")
// 	var responseMessage message
	
// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		w.WriteHeader(http.StatusNoContent)
// 		responseMessage.Message = "Kindly enter data with the user account"
// 		json.NewEncoder(w).Encode(responseMessage)
// 	}
	
// 	json.Unmarshal(reqBody, &user)
// 	newUser := account {
// 		Id: primitive.NewObjectID(),
// 		Email: user.Email,
// 		Pass: user.Pass,
// 	}
// 	result, err := userCollection.InsertOne(ctx, newUser)
// 	if err != nil {
// 		w.WriteHeader(http.StatusNoContent)
// 		responseMessage.Message = "Error to insert data"
// 		json.NewEncoder(w).Encode(responseMessage)
// 	}
// 	fmt.Printf("%v", result)
// 	w.WriteHeader(http.StatusCreated)
// 	fmt.Fprintf(w, "Account created successfully")
// }

// func getOneUser(w http.ResponseWriter, r *http.Request) {
// 	// var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "usuarios")
// 	w.Header().Set("Content-Type", "application/json")
// 	var responseMessage message
// 	var checkAccount account
// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		w.WriteHeader(http.StatusNoContent)
// 		responseMessage.Message = "Kindly enter data with the user account"
// 		json.NewEncoder(w).Encode(responseMessage)
// 	}

// 	json.Unmarshal(reqBody, &checkAccount)
// 	for _, singleAccount := range accounts {
// 		if singleAccount.Email == checkAccount.Email {
// 			if singleAccount.Pass == checkAccount.Pass {
// 				w.WriteHeader(http.StatusAccepted)
// 				responseMessage.Message = "Access granted"
// 				json.NewEncoder(w).Encode(responseMessage)
// 				return
// 			} else {
// 				w.WriteHeader(http.StatusNotAcceptable)
// 				responseMessage.Message = "Wrong password"
// 				json.NewEncoder(w).Encode(responseMessage)
// 				return
// 			}
// 		}
// 	}
// 	w.WriteHeader(http.StatusNotFound)
// 	responseMessage.Message = "User not found"
// 	json.NewEncoder(w).Encode(responseMessage)
// }

// func getAllUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(accounts)
// }

func callMemsim(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var responseMessage models.Message
	var configMemsim models.Memsim
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		responseMessage.Message = "Kindly enter data for the memsim command"
		json.NewEncoder(w).Encode(responseMessage)
	}

	json.Unmarshal(reqBody, &configMemsim)
	fmt.Printf("%v", configMemsim)

	var result = execMemsim(configMemsim.Ciclos, configMemsim.Unidades)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func execMemsim(ciclos int, unidades []int) models.MemsimRes {
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
	return result
}

func work(proceso int, unidad int, tam int) {
	var inicioProceso = "El proceso # " + strconv.Itoa(proceso) + ", empezó a trabajar con la unidad: '" + strconv.Itoa(unidad) + "'"
	procesos = append(procesos, inicioProceso)
	
	time.Sleep(time.Duration(tam) * time.Millisecond)
	
	var finProceso = "El proceso # " + strconv.Itoa(proceso) + ", terminó de trabajar con la unidad: '" + strconv.Itoa(unidad) + "'"
	procesos = append(procesos, finProceso)
}
