package main

import (
	"fmt"
	"sync"
	"time"
)

func main() { // Main tiene una rutina asignada por defecto
	now := time.Now()
	var wg sync.WaitGroup // Declarando nuestro wait group
	wg.Add(1)             // Indicamos la cantidad de rutinas a esperar
	/*
		En lugar de llamar a go work utilizamos una funcion anonima
	*/

	go func() {
		defer wg.Done() // Mensaje region critica
		work()
	}() // ----------> FORK

	wg.Wait() // JOIN <----------
	fmt.Println("Ha transcurrido: ", time.Since(now))
	fmt.Println("La rutina principal ha terminado")
}

func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Rutina work esta trabajando")
}
