package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"

	"./procesos"
)

func Aumentador(p procesos.Proceso) {
	for {
		p.Tiempo += 1
		p.Mostrar()
		time.Sleep(time.Millisecond * 500)
	}
}

func cliente() { //conexion
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	var msgd procesos.Proceso //guardo el proceso
	err = gob.NewDecoder(c).Decode(&msgd)
	if err != nil {
		fmt.Println(err)
	}
	go Aumentador(msgd)
	defer fin(c, msgd)

}

func fin(c net.Conn, msgd procesos.Proceso) {
	fmt.Println("Regresando al servidor")
	msg := msgd
	err := gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

func main() {
	go cliente()

	var input string
	fmt.Scanln(&input)
}
