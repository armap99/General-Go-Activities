package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"net/rpc"
	"os"
	"time"

	"./claseschat"
)

func obtenercadenaespacios() string {
	s := ""
	for s == "" {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		s = scanner.Text()
	}
	return s
}

func conexionesActuales(){
	var result string
	//Tcp
	conexionRCP, err := rpc.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}
	err = conexionRCP.Call("Server.RegresarServidores",1,&result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}

func conexionTCP() string{
	regreso := "8000"
	var result string
	var puertorcp string
	//Tcp
	conexionRCP, err := rpc.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}
	err = conexionRCP.Call("Server.RegresarServidores",1,&result)
	if err != nil {
		fmt.Println(err)
	} else {
		var tematica string
		fmt.Println(result)
		fmt.Println("\n")
		fmt.Println("Ingrese al chat que quiere acceder: ")
		fmt.Scan(&tematica)
		err = conexionRCP.Call("Server.ObtenerPuertoPorTema",tematica,&puertorcp)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(puertorcp)
			fmt.Println("\n")
			regreso = puertorcp
		}
	}
	return regreso
}

func conexion() net.Conn {
	puerto := conexionTCP()
	//rcp
	c, err := net.Dial("tcp", ":"+puerto)
	if err != nil {
		fmt.Println(err)
	}
	return c
}

func CrearUsuario() *claseschat.Usuario {
	usuario := new(claseschat.Usuario)
	var nombre string
	fmt.Println("Antes de comenzar ingresa tu Nombre: ")
	fmt.Scan(&nombre)
	usuario.Nombre = nombre
	usuario.Conectado = 1

	return usuario
}

func regresarUsuario(c net.Conn, usuario *claseschat.Usuario) {
	msg := usuario
	err := gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}

}

func MenuUsuario(c net.Conn, usuario *claseschat.Usuario) {
	var opcliente int
	if usuario.Nombre == "Admin" {
		for {
			fmt.Println("Opciones cliente: ")
			fmt.Println("1-Enviar mensaje de texto ")
			fmt.Println("2-Ver salas ")
			fmt.Println("6-Salir ")
			fmt.Scan(&opcliente)
			if opcliente == 1 {
				mensaje := new(claseschat.Mensaje)
				mensaje.Enviador = usuario.Nombre
				mensaje.DiaEnvio = time.Now()
				destinatario := "Todos"
				var contendio string
				mensaje.Destinatario = destinatario
				fmt.Println("Texto: ")
				contendio = obtenercadenaespacios()
				mensaje.Contenido = contendio
				MandarMensaje(c, mensaje)
	
			}  else if opcliente == 6 {
				return
			} else if opcliente == 2 {
				conexionesActuales()
			}
		}

	}else{
		for {
			fmt.Println("Opciones cliente: ")
			fmt.Println("1-Enviar mensaje de texto ")
			fmt.Println("6-Salir ")
			fmt.Scan(&opcliente)
			if opcliente == 1 {
				mensaje := new(claseschat.Mensaje)
				mensaje.Enviador = usuario.Nombre
				mensaje.DiaEnvio = time.Now()
				destinatario := "Todos"
				var contendio string
				mensaje.Destinatario = destinatario
				fmt.Println("Texto: ")
				contendio = obtenercadenaespacios()
				mensaje.Contenido = contendio
				MandarMensaje(c, mensaje)
	
			}  else if opcliente == 6 {
				return
			}
		}
	}
	

}

func MandarMensaje(conexion net.Conn, mensaje *claseschat.Mensaje) {
	err := gob.NewEncoder(conexion).Encode(mensaje)
	if err != nil {
		fmt.Println(err)
	}

}

func EsperandoMensajes(c net.Conn, usuario *claseschat.Usuario) {
	for {
		var msgs claseschat.Mensaje
		err := gob.NewDecoder(c).Decode(&msgs)
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			usuario.MensajesRecibidos = append(usuario.MensajesRecibidos, msgs)
			msgs.MostrarMensajeRecibidos()
		}

	}
}

func main() {
	conexion := conexion()
	usuario := CrearUsuario()
	regresarUsuario(conexion, usuario)
	go EsperandoMensajes(conexion, usuario)
	MenuUsuario(conexion, usuario)

	var input string
	fmt.Scanln(&input)

}
