package main

import (
	"fmt"
	"io"
	"net/http"
)

type Response struct {
	Source string
	Result string
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/logup", logup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/memsim", memsim)
	http.ListenAndServe(":8080", nil)
}

func root(rw http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(rw, "{message: \"Sopes 2\"}")
}

func logup(rw http.ResponseWriter, r *http.Request) {

}

func login(rw http.ResponseWriter, r *http.Request) {

}

func memsim(rw http.ResponseWriter, r *http.Request) {

}