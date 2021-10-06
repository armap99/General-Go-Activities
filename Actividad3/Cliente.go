package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"

	"./procesos"
)

func Aumentador(p *procesos.Proceso) {
	for {
		p.Tiempo += 1
		p.Mostrar()
		time.Sleep(time.Millisecond * 500)
	}
}

func fin(c net.Conn, msgd *procesos.Proceso) {
	fmt.Println("Regresando al servidor")
	msg := msgd
	err := gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

func conexion() net.Conn {
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}
	return c
}

func Obtener(c net.Conn) *procesos.Proceso {
	msgd := new(procesos.Proceso) //guardo el proceso
	//var msgd procesos.Proceso
	err := gob.NewDecoder(c).Decode(&msgd)
	if err != nil {
		fmt.Println(err)
	}
	return msgd
}

func main() {
	conexion := conexion()
	prose := Obtener(conexion)
	go Aumentador(prose)
	var input string
	fmt.Scanln(&input)
	fin(conexion, prose)
}
