package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type message struct {
	Message string `json:"message"`
}

type account struct {
	Email string `json:"email"`
	Pass string `json:"pass"`
}

type memsim struct {
	Ciclos int `json:"ciclos"`
	Unidades []int `json:"unidades"`
}

type allAccount []account

var accounts = allAccount{
	{
		Email: "usuario@example.com",
		Pass: "super-password",
	},
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	// router.Headers("Content-Type", "application/json")
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/logup", createAccount)
	router.HandleFunc("/login", getOneAccount)
	router.HandleFunc("/users", getAllAccount)
	router.HandleFunc("/memsim", callMemsim)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var responseMessage message
	responseMessage.Message ="Welcome home!"
	json.NewEncoder(w).Encode(responseMessage)
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var responseMessage message
	var newAccount account
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		responseMessage.Message = "Kindly enter data with the user account"
		json.NewEncoder(w).Encode(responseMessage)
	}

	json.Unmarshal(reqBody, &newAccount)
	accounts = append(accounts, newAccount)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Account created successfully")
}

func getOneAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var responseMessage message
	var checkAccount account
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		responseMessage.Message = "Kindly enter data with the user account"
		json.NewEncoder(w).Encode(responseMessage)
	}

	json.Unmarshal(reqBody, &checkAccount)
	for _, singleAccount := range accounts {
		if singleAccount.Email == checkAccount.Email {
			if singleAccount.Pass == checkAccount.Pass {
				w.WriteHeader(http.StatusAccepted)
				responseMessage.Message = "Access granted"
				json.NewEncoder(w).Encode(responseMessage)
				return
			} else {
				w.WriteHeader(http.StatusNotAcceptable)
				responseMessage.Message = "Wrong password"
				json.NewEncoder(w).Encode(responseMessage)
				return
			}
		}
	}
	w.WriteHeader(http.StatusNotFound)
	responseMessage.Message = "User not found"
	json.NewEncoder(w).Encode(responseMessage)
}

func getAllAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accounts)
}

func callMemsim(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var responseMessage message
	var configMemsim memsim
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		responseMessage.Message = "Kindly enter data for the memsim command"
		json.NewEncoder(w).Encode(responseMessage)
	}

	json.Unmarshal(reqBody, &configMemsim)
	fmt.Printf("%v", configMemsim)

	execMemsim(configMemsim.Ciclos, configMemsim.Unidades)

	w.WriteHeader(http.StatusOK)
	responseMessage.Message = "Response OK"
	json.NewEncoder(w).Encode(responseMessage)
}

func execMemsim(ciclos int, unidades []int) {
	now := time.Now()
	var process int = 0
	fmt.Println("::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::")
	for i := 1; i <= ciclos; i++ {
		var wg sync.WaitGroup // Declarando nuestro wait group
		fmt.Println("	::::::::::::	Ciclo de trabajo: ", i, "	::::::::::::")
		fmt.Println("::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::")
		for _, unidad := range unidades {
			wg.Add(1) // Indicamos la cantidad de rutinas a esperar
			go func() {
				defer wg.Done() // Mensaje region critica
				process += 1
				work(process, unidad, len(unidades))
			}()
		}
		wg.Wait()
		fmt.Println("::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::")
	}
	// logStatus[userName][strconv.Itoa(process)] =unidades

	fmt.Println("Ha transcurrido: ", time.Since(now))
	fmt.Println("La rutina principal ha terminado")
}

func work(proceso int, unidad int, tam int) {
	fmt.Println("| ⌚ El proceso 💼 # ", proceso, ", empezó a trabajar con la unidad: '", strconv.Itoa(unidad), "' |")
	time.Sleep(time.Duration(tam) * time.Millisecond)
	fmt.Println("| ✅ El proceso 💼 # ", proceso, ", terminó de trabajar con la unidad: '", strconv.Itoa(unidad), "' |")
}
