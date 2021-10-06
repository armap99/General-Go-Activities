package main

import (
	"encoding/gob"
	"fmt"
	"net"

	"./procesos"
)

func servidor(lprocesos *procesos.Bloque) {
	servidor, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		cliente, err := servidor.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		msg := lprocesos.Primero()
		err = gob.NewEncoder(cliente).Encode(msg)
		lprocesos.RemoveIndex(0)
		go handleClient(cliente, lprocesos)
	}
}

func handleClient(c net.Conn, proseos *procesos.Bloque) {
	var msg procesos.Proceso
	err := gob.NewDecoder(c).Decode(&msg)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		//fmt.Println("entre")
		proseos.Procesos = append(proseos.Procesos, msg)
	}

}

func main() {
	lprocesos := new(procesos.Bloque)
	/*lprocesos := procesos.Bloque{
		Procesos: []procesos.Proceso{
			{Id: 0, Tiempo: 0, Estatus: 1},
			{Id: 1, Tiempo: 0, Estatus: 1},
			{Id: 2, Tiempo: 0, Estatus: 1},
			{Id: 3, Tiempo: 0, Estatus: 1},
			{Id: 4, Tiempo: 0, Estatus: 1},
		},
	}*/
	s := []procesos.Proceso{
		{Id: 0, Tiempo: 0, Estatus: 1},
		{Id: 1, Tiempo: 0, Estatus: 1},
		{Id: 2, Tiempo: 0, Estatus: 1},
		{Id: 3, Tiempo: 0, Estatus: 1},
		{Id: 4, Tiempo: 0, Estatus: 1},
	}
	lprocesos.Procesos = append(lprocesos.Procesos, s...)
	go lprocesos.Start()
	go servidor(lprocesos)

	var input string
	fmt.Scanln(&input)
}
