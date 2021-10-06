package main

import (
	"fmt"

	"./procesos"
)

func main() {
	auxid := 0
	foraux := true
	lProcesos := new(procesos.Bloque)
	var c chan string = make(chan string)
	for foraux {
		var op int
		fmt.Println("Bienvenido")
		fmt.Println("1-Agregar proceso")
		fmt.Println("2-Mostrar procesos")
		fmt.Println("3-Terminar proceso")
		fmt.Println("4-Salir")
		fmt.Scan(&op)
		if op == 1 {
			proceso := procesos.Proceso{Id: uint64(auxid), Tiempo: 0, Estatus: 1}
			lProcesos.Procesos = append(lProcesos.Procesos, proceso)
			auxid = auxid + 1
			lProcesos.Start()

		} else if op == 2 {
			lProcesos.MostrarGeneral(c)

		} else if op == 3 {
			var idter uint64
			fmt.Println("Ingrese el id que desea eliminar: ")
			fmt.Scan(&idter)
			lProcesos.Eliminar(idter)

		} else if op == 4 {
			foraux = false
		} else {
			lProcesos.Stop(c)
		}

	}
}
