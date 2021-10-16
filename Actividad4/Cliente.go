package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
)

type StudentDataRecive struct {
	Name    string
	Subject string
	Score   float64
}

func client() {
	c, err := rpc.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var continuar bool = true
	for continuar {
		switch mainMenu() {
		case 1:
			var name string
			var subject string
			var score float64
			var result string

			fmt.Println("Agregar nuevo registro")
			fmt.Print("Nombre: ")
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				name = scanner.Text()
			} else {
				fmt.Println("Error al escanear...")
				os.Exit(3)
			}
			fmt.Print("Materia: ")
			if scanner.Scan() {
				subject = scanner.Text()
			} else {
				fmt.Println("Error al escanear...")
				os.Exit(3)
			}
			fmt.Print("Calificacion: ")
			fmt.Scanln(&score)

			data := StudentDataRecive{
				Name:    name,
				Subject: subject,
				Score:   score,
			}

			err = c.Call("Server.AddStudentData", data, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(result)
				fmt.Println("\n\n")
			}
			break
		case 2:
			var name string
			var result string

			fmt.Println("Promedio del alumno \n")
			fmt.Print("Nombre del alumno: ")
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				name = scanner.Text()
			} else {
				fmt.Println("Error al escanear...")
				os.Exit(3)
			}
			err = c.Call("Server.GetStudentAverage", name, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(result)
				fmt.Println("\n\n")
			}
			break
		case 3:
			var petition = "petition"
			var result string
			err = c.Call("Server.GetGeneralAverageByStudents", petition, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(result)
				fmt.Println("\n\n")
			}
			break
		case 4:
			var subject string
			var result string

			fmt.Println("Promedio de una materia: ")
			fmt.Print("Materia: ")
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				subject = scanner.Text()
			} else {
				fmt.Println("Error al escanear...")
				os.Exit(3)
			}

			err = c.Call("Server.GetAverageBySubject", subject, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(result)
				fmt.Println("\n\n")
			}
			break
		case 5:
			continuar = false
			break
		default:
			break
		}
	}
}

func mainMenu() int64 {
	var option int64

	fmt.Println("1-Agregar calificacion de materia por alumno")
	fmt.Println("2-Mostrar promedio de alumno")
	fmt.Println("3-Mostrar promedio general")
	fmt.Println("4-Mostrar promedio por materia")
	fmt.Println("5-Salir")
	fmt.Print("Seleccione la opcion: ")
	fmt.Scanln(&option)

	return option
}

func main() {
	client()
}
