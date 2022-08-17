package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

var userName string

func main() {
	// MAIN MENU
	menu := make(map[string]string)
	menu["1"] = "Nueva ejecución"
	menu["2"] = "Reporte"
	menu["3"] = "Salir"

	for {
		clearConsole()

		// WELCOME MESSAGE
		b, err := ioutil.ReadFile("welcome.txt")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))

		// MAIN MENU
		for k, v := range menu {
			println(k, v)
		}

		var choice string
		println("Ingrese su opción: ")
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			newExec()
		case "2":
			println("Reporte")
		case "3":
			println("Salir")
			return
		default:
			println("Opción no válida")
		}
	}
}

func newExec() {
	clearConsole()

	println("**************************************************************************")
	println("***                         NUEVA EJECUCIÓN                            ***")
	println("**************************************************************************")

	menu := make(map[string]string)
	menu["1"] = "Ingresar nombre"
	menu["2"] = "Regresar"
	for {
		for k, v := range menu {
			println(k, v)
		}

		var choice string
		println("Ingrese su opción: ")
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			println("Ingresar nombre")
			fmt.Scanln(&userName)
			selectCommands()
		case "2":
			println("Salir")
			return
		default:
			println("Opción no válida")
		}
	}
}

func selectCommands() {
	println("\n**************************************************************************\n")
	// COMMAND MENU
	menu := make(map[string]string)
	menu["1"] = "IOTOP"
	menu["2"] = "TOP"
	menu["3"] = "Regresar"
	for {
		for k, v := range menu {
			println(k, v)
		}

		var choice string
		println("Ingrese su opción: ")
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			cmdIOTOP()
		case "2":
			println("TOP")
		case "3":
			println("Salir")
			userName = ""
			clearConsole()
			return
		default:
			println("Opción no válida")
		}
	}
}

func cmdIOTOP() {
	var answer string
	for {
		clearConsole()

		println("**************************************************************************")
		println("***                              IOTOP                                 ***")
		println("**************************************************************************")
		fmt.Println("*** USUARIO: ", userName)
		println("**************************************************************************\n")

		out, err := exec.Command("iotop", "-b", "-n1").Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(out))

		println("¿Ejecutar de nuevo? (y/n): ")
		fmt.Scanln(&answer)
		if answer == "n" || answer == "N" {
			clearConsole()
			break
		} else if answer == "y" || answer == "Y" {
			continue
		} else {
			println("Respuesta no válida.")
		}
	}
}

func cmdTOP() {
	var answer string
	for {
		clearConsole()

		println("**************************************************************************")
		println("***                              TOP                                 ***")
		println("**************************************************************************")
		fmt.Println("*** USUARIO: ", userName)
		println("**************************************************************************\n")

		out, err := exec.Command("sudo insmod modules/proc_mod.ko").Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(out))

		println("¿Ejecutar de nuevo? (y/n): ")
		fmt.Scanln(&answer)
		if answer == "n" || answer == "N" {
			clearConsole()
			break
		} else if answer == "y" || answer == "Y" {
			continue
		} else {
			println("Respuesta no válida.")
		}
	}

}

func clearConsole() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
