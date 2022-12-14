package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

var userName string
var logMap = map[string]map[string]int{}
var logStatus = map[string]map[string]string{}

func main() {
	// MAIN MENU
	menu := make(map[string]string)
	menu["1"] = "Nueva ejecuci贸n"
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
		println("Ingrese su opci贸n: ")
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			newExec()
		case "2":
			reporte()
		case "3":
			println("See you later 馃槂!!!")
			return
		default:
			println("Opci贸n no v谩lida 馃ジ")
		}
	}
}

func newExec() {
	clearConsole()

	println("**************************************************************************")
	println("***                         NUEVA EJECUCI脫N                            ***")
	println("**************************************************************************")

	menu := make(map[string]string)
	menu["1"] = "Ingresar nombre"
	menu["2"] = "Regresar"
	for {
		for k, v := range menu {
			println(k, v)
		}

		var choice string
		println("Ingrese su opci贸n: ")
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			println("Ingresar nombre")
			fmt.Scanln(&userName)
			logMap[userName] = map[string]int{}
			logStatus[userName] = map[string]string{}
			selectCommands()
			userName = ""
		case "2":
			println("Salir")
			return
		default:
			println("Opci贸n no v谩lida 馃ジ")
		}
	}
}

func reporte() {
	clearConsole()

	println("**************************************************************************")
	println("***                             REPORTES                               ***")
	println("**************************************************************************")

	menu := make(map[string]string)
	menu["1"] = "Bit谩cora"
	menu["2"] = "Estado Simulacion"
	menu["3"] = "Regresar"
	for {
		for k, v := range menu {
			println(k, v)
		}

		var choice string
		println("Ingrese su opci贸n: ")
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			// BITACORA
			bitacora()
		case "2":
			bitacoraStatus()
		case "3":
			println("Salir")
			return
		default:
			println("Opci贸n no v谩lida 馃ジ")
		}
	}
}

func bitacora() {
	println("**************************************************************************")
	println("***                            BITACORA                                ***")
	println("**************************************************************************")
	cadena := "[\n"
	for k, v := range logMap {
		cadena += "\t{\n"
		cadena += "\t\t\"usuario\": " + k + "\n"
		cadena += "\t\t\"actividad\": [\n"
		for l, w := range v {
			cadena += "\t\t\t{\n"
			cadena += "\t\t\t\t\"funcion\": " + l + ",\n"
			cadena += "\t\t\t\t\"ejecutandose\": " + strconv.Itoa(w) + ",\n"
			cadena += "\t\t\t},\n"
		}
		cadena += "\t\t]\n"
		cadena += "\t},\n"
	}
	cadena += "]\n"

	b := []byte(cadena)
	err := ioutil.WriteFile("bitacora.json", b, 0644)
	if err != nil {
		log.Fatal(err)
	} else {
		println("Reporte de bitacora generado con exito.\nEjecutar cat bitacora.json para verlo")
	}
}

func bitacoraStatus() {
	println("**************************************************************************")
	println("***                      Estado de Simulacion                          ***")
	println("**************************************************************************")
	cadena := "[\n"
	for k, v := range logStatus {
		cadena += "\t{\n"
		cadena += "\t\t\"usuario\": " + k + "\n"
		cadena += "\t\t\"ejecucion\": [\n"
		for l, w := range v {
			cadena += "\t\t\t{\n"
			cadena += "\t\t\t\t\"procesos\": " + l + ",\n"
			cadena += "\t\t\t\t\"unidades\": " + w + ",\n"
			cadena += "\t\t\t},\n"
		}
		cadena += "\t\t]\n"
		cadena += "\t},\n"
	}
	cadena += "]\n"

	b := []byte(cadena)
	err := ioutil.WriteFile("bitacoraSimulacion.json", b, 0644)
	if err != nil {
		log.Fatal(err)
	} else {
		println("Reporte de bitacora generado con exito.\nEjecutar cat bitacora.json para verlo")
	}
}

func selectCommands() {
	println("\n**************************************************************************\n")
	fmt.Println(userName)
	// COMMAND MENU
	menu := make(map[string]string)
	menu["1"] = "IOTOP"
	menu["2"] = "TOP"
	menu["3"] = "STRACE"
	menu["4"] = "MEMSIM"
	menu["5"] = "Regresar"
	for {
		for k, v := range menu {
			println(k, v)
		}

		var choice string
		println("Ingrese su opci贸n: ")
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			cmdIOTOP()
		case "2":
			cmdTOP()
		case "3":
			cmdSTRACE()
		case "4":
			cmdMEMSIM()
		case "5":
			println("Salir")
			clearConsole()
			return
		default:
			println("Opci贸n no v谩lida 馃ジ")
		}
	}
}

func cmdIOTOP() {
	var answer string
	for {
		clearConsole()
		logMap[userName]["IOTOP"]++

		println("**************************************************************************")
		println("***                              IOTOP                                 ***")
		println("**************************************************************************")
		fmt.Println("*** 				USUARIO: ", userName, "				***")
		println("**************************************************************************\n")

		out, err := exec.Command("iotop", "-b", "-n1").Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(out))

		println("驴Ejecutar de nuevo? (y/n): ")
		fmt.Scanln(&answer)
		if answer == "n" || answer == "N" {
			clearConsole()
			break
		} else if answer == "y" || answer == "Y" {
			continue
		} else {
			println("Respuesta no v谩lida 馃ジ")
		}
	}
}

func cmdTOP() {
	var answer string
	for {
		clearConsole()
		logMap[userName]["TOP"]++

		println("**************************************************************************")
		println("***                              TOP                                 ***")
		println("**************************************************************************")
		fmt.Println("*** 				USUARIO: ", userName, "				***")
		println("**************************************************************************\n")

		// out, err := exec.Command("sudo insmod modules/proc_mod.ko").Output()
		out, err := exec.Command("/usr/bin/cat", "/proc/proc_mod").Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(out))

		println("驴Ejecutar de nuevo? (y/n): ")
		fmt.Scanln(&answer)
		if answer == "n" || answer == "N" {
			clearConsole()
			break
		} else if answer == "y" || answer == "Y" {
			continue
		} else {
			println("Respuesta no v谩lida 馃ジ")
		}
	}
}

func cmdSTRACE() {
	var answer string
	for {
		clearConsole()
		logMap[userName]["STRACE"]++

		for {
			println("Ingrese un comando: ")
			com := bufio.NewScanner(os.Stdin)
			if com.Scan() {
				println("**************************************************************************")
				println("***                              STRACE SYSTEM                         ***")
				println("**************************************************************************")
				fmt.Println("*** 				USUARIO: ", userName, "				***")
				println("**************************************************************************\n")
				strace(strings.Fields(com.Text()))
				break
			}
		}

		println("驴Ejecutar de nuevo? (y/n): ")
		fmt.Scanln(&answer)
		if answer == "n" || answer == "N" {
			clearConsole()
			break
		} else if answer == "y" || answer == "Y" {
			continue
		} else {
			println("Respuesta no v谩lida 馃ジ")
		}
	}
}

func cmdMEMSIM() {
	var answer string
	for {
		clearConsole()
		logMap[userName]["MEMSIM"]++

		println("**************************************************************************")
		println("***                            MEMORY SIMULATION                       ***")
		println("**************************************************************************")
		fmt.Println("*** 				USUARIO: ", userName, "				***")
		println("**************************************************************************\n")

		println("Ingrese la cantidad de ciclos de trabajo (solo un entero): ")
		var cycles int
		fmt.Scanln(&cycles)
		if cycles <= 0 {
			println("Cantidad de ciclos no v谩lida.")
			continue
		}
		println("Ingrese las unidades de memoria (separados por coma): ")
		var memUnits string
		fmt.Scanln(&memUnits)
		memsim(cycles, memUnits)

		println("驴Ejecutar de nuevo? (y/n): ")
		fmt.Scanln(&answer)
		if answer == "n" || answer == "N" {
			clearConsole()
			break
		} else if answer == "y" || answer == "Y" {
			continue
		} else {
			println("Respuesta no v谩lida 馃ジ")
		}
	}
}

func clearConsole() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func strace(command []string) {
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

func memsim(ciclos int, unidades string) {
	
	units_arr := strings.Split(unidades, ",")
	size := len(units_arr)
	now := time.Now()
	var process int = 0
	fmt.Println("::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::")
	for i := 1; i <= ciclos; i++ {
		var wg sync.WaitGroup // Declarando nuestro wait group
		fmt.Println("	::::::::::::	Ciclo de trabajo: ", i, "	::::::::::::")
		fmt.Println("::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::")
		for j := size - 1; j >= 0; j-- {
			wg.Add(1) // Indicamos la cantidad de rutinas a esperar
			value := units_arr[j]
			go func() {
				defer wg.Done() // Mensaje region critica
				process = process + 1
				work(process, value, size)
			}()
		}
		wg.Wait()
		fmt.Println("::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::")
	}
	logStatus[userName][strconv.Itoa(process)] =unidades

	fmt.Println("Ha transcurrido: ", time.Since(now))
	fmt.Println("La rutina principal ha terminado")
}

func work(proceso int, unidad string, tam int) {
	fmt.Println("| 鈱? El proceso 馃捈 # ", proceso, ", empez贸 a trabajar con la unidad: '", unidad, "' |")
	time.Sleep(time.Duration(tam) * time.Millisecond)
	fmt.Println("| 鉁? El proceso 馃捈 # ", proceso, ", termin贸 de trabajar con la unidad: '", unidad, "' |")
}
