package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"syscall"
	"bufio"
	"strings"
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
	menu["3"] = "STRACE"
	menu["4"] = "Regresar"
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
			cmdTOP()
		case "3":	
			
			cmdSTRACE()
		case "4":
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
		fmt.Println("*** 						USUARIO: 								  ***", userName)
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
		fmt.Println("*** 						USUARIO: 								  ***", userName)
		println("**************************************************************************\n")

		// out, err := exec.Command("sudo insmod modules/proc_mod.ko").Output()
		out, err := exec.Command("/usr/bin/cat", "/proc/proc_mod").Output()
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

func cmdSTRACE() {
		var answer string
	
	for {
		for {
			println("***             INGRESE UN COMANDO                       ***")
			com := bufio.NewScanner(os.Stdin)
			if com.Scan() {
				println("**************************************************************************")
				println("***                              STRACE SYSTEM                         ***")
				println("**************************************************************************")
				fmt.Println("*** 						USUARIO: 					  ***", userName)
				println("**************************************************************************\n")
				strace(strings.Fields(com.Text()))
				break;
			}
		}
		
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

func strace(command []string){
	var regs syscall.PtraceRegs
	var ss syscallCounter

	ss = ss.init()

	cmd := exec.Command(command[0], command[1])
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Ptrace: true,
	}

	cmd.Start()
	err := cmd.Wait()
	if err != nil {
		fmt.Printf("Wait returned: %v\n", err)
	}

	pid := cmd.Process.Pid
	exit := true

	for {
		if exit {
			err = syscall.PtraceGetRegs(pid, &regs)
			if err != nil {
				break
			}
			ss.inc(regs.Orig_rax)
		}

		err = syscall.PtraceSyscall(pid, 0)
		if err != nil {
			panic(err)
		}

		_, err = syscall.Wait4(pid, nil, 0, nil)
		if err != nil {
			panic(err)
		}

		exit = !exit
	}

	ss.print()
}
