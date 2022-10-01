package main

import (
	"fmt"
	"time"
)

func main() { // Main tiene una rutina asignada por defecto
	fmt.Println("Hello, World!")
	// fmt.Println("Siuuu")
	now := time.Now()
	task1()
	task2()
	task3()
	task4()
	time.Sleep(time.Second)
	fmt.Println("Ha transcurrido: ", time.Since(now))
}

func task1() {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Task 1")
}

func task2() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Task 2")
}

func task3() {
	fmt.Println("Task 3")
}

func task4() {
	time.Sleep(300 * time.Millisecond)
	fmt.Println("Task 4")
}
