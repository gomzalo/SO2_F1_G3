package main

import (
	"encoding/json"
	"log"
	"net/http"
	"server/configs"
	"server/controllers"
	"server/models"

	"github.com/gorilla/mux"
)

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			// Just put some headers to allow CORS...
      		w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
			w.Header().Add("Access-Control-Allow-Credentials", "true")
			w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("content-type", "application/json;charset=UTF-8")
			if req.Method == "OPTIONS" {
				w.WriteHeader(http.StatusNoContent)
				return
			}
			// and call next handler!
			next.ServeHTTP(w, req)
		})
}

func main() {
	configs.ConnectDB()
	
	router := mux.NewRouter().StrictSlash(true)
	router.Use(middlewareCors)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/logup", controllers.CreateUser)
	router.HandleFunc("/login", controllers.GetOneUser)
	router.HandleFunc("/users", controllers.GetAllUser)
	router.HandleFunc("/memsim", controllers.CreateMemsim)
	router.HandleFunc("/memsim/user", controllers.GetAllMemsim)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var responseMessage models.Message
	responseMessage.Message ="Welcome home!"
	json.NewEncoder(w).Encode(responseMessage)
}

// func callMemsim(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var responseMessage models.Message
// 	var configMemsim models.Memsim
// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		w.WriteHeader(http.StatusNoContent)
// 		responseMessage.Message = "Kindly enter data for the memsim command"
// 		json.NewEncoder(w).Encode(responseMessage)
// 	}

// 	json.Unmarshal(reqBody, &configMemsim)
// 	fmt.Printf("%v", configMemsim)

// 	var result = execMemsim(configMemsim.Ciclos, configMemsim.Unidades)
// 	result.Id = primitive.NewObjectID()
// 	result.UserId = configMemsim.UserId

// 	controllers.CreateMemsim(result)

// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(result)
// }

// func execMemsim(ciclos int, unidades []int) models.MemsimRes {
// 	var result models.MemsimRes
// 	size := len(unidades)
// 	now := time.Now()
// 	var process int = 0

// 	for i := 1; i <= ciclos; i++ {
// 		var proceso models.MemsimProc
// 		var wg sync.WaitGroup // Declarando nuestro wait group
		
// 		proceso.Ciclo = i
		
// 		for j := size - 1; j >= 0; j-- {
// 			wg.Add(1) // Indicamos la cantidad de rutinas a esperar
// 			value := unidades[j]
// 			go func() {
// 				defer wg.Done() // Mensaje region critica
// 				process = process + 1
// 				work(process, value, size)
// 			}()
// 		}
// 		wg.Wait()
// 		proceso.Procesos = procesos
// 		result.Memsim = append(result.Memsim, proceso)
// 		procesos = nil
// 	}

// 	var duracion = time.Since(now).Milliseconds()
// 	result.Duracion = duracion
// 	result.Procesos = process
// 	result.Unidades = unidades

// 	return result
// }

// func work(proceso int, unidad int, tam int) {
// 	var inicioProceso = "El proceso # " + strconv.Itoa(proceso) + ", empezó a trabajar con la unidad: '" + strconv.Itoa(unidad) + "'"
// 	procesos = append(procesos, inicioProceso)
	
// 	time.Sleep(time.Duration(tam) * time.Millisecond)
	
// 	var finProceso = "El proceso # " + strconv.Itoa(proceso) + ", terminó de trabajar con la unidad: '" + strconv.Itoa(unidad) + "'"
// 	procesos = append(procesos, finProceso)
// }
