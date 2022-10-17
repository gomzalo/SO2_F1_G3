package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/logup", logup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/memsim", memsim)
	http.ListenAndServe(":8080", nil)
}

func root(rw http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"status": true,
		"message": "Sopes 2",
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("got / request\n")
		io.WriteString(rw, err.Error())
	}
	fmt.Printf("got / request\n")
	io.WriteString(rw, string(jsonData))
}

func logup(rw http.ResponseWriter, r *http.Request) {

}

func login(rw http.ResponseWriter, r *http.Request) {

}

func memsim(rw http.ResponseWriter, r *http.Request) {

}