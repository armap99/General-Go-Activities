package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"net/rpc"

	"./claseschat"
)

func AsigancionDePuerto() string {
	conexionRCP, err := rpc.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}

	var result string
	err = conexionRCP.Call("Server.ObtenerPuertoLibre",1,&result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Puerto asigando: ",result)
		fmt.Println("\n\n")
	}
	return result
}

func AgregarAInterme(dta *claseschat.Servidor )  {
	conexionRCP, err := rpc.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}
	var result string
	err = conexionRCP.Call("Server.AgregarSala",dta,&result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
func ActualizarServidores()  {
	conexionRCP, err := rpc.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}
	var result string
	err = conexionRCP.Call("Server.ImprimirServidores",1,&result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func servidor(lservidor *claseschat.Servidor) {
	result := AsigancionDePuerto()
	lservidor.Puerto = result
	AgregarAInterme(lservidor)
	servidor, err := net.Listen("tcp", ":" + result)
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
		ActualizarServidores()
		go handleClient(cliente, lservidor)
	}
}

func handleClient(c net.Conn, servidor *claseschat.Servidor) { // se agrega al usuario al servidor
	var msg claseschat.Usuario
	err := gob.NewDecoder(c).Decode(&msg)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		msg.Conexion = c
		servidor.Usuarios = append(servidor.Usuarios, msg)
	}
	for {
		mensaje := new(claseschat.Mensaje)
		err := gob.NewDecoder(c).Decode(&mensaje)
		if err != nil {
			fmt.Println(err)
			return
		} else {

			fmt.Println("Mensaje enviado: [ ", mensaje.Enviador, " | ", mensaje.Destinatario, " ] ")
			if mensaje.Destinatario != "Todos" {
				EnviarMensaje(mensaje, servidor)
			} else {
				EnviarMensajeGeneral(mensaje, servidor)
			}
		}
	}

}

func EnviarMensaje(msg *claseschat.Mensaje, servidor *claseschat.Servidor) {
	var aux net.Conn
	for i := 0; i < len(servidor.Usuarios); i++ {
		if servidor.Usuarios[i].Nombre == msg.Destinatario {
			aux = servidor.Usuarios[i].Conexion
			servidor.Usuarios[i].MensajesRecibidos = append(servidor.Usuarios[i].MensajesRecibidos, *msg)
		}
	}
	err := gob.NewEncoder(aux).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
}

func EnviarMensajeGeneral(msg *claseschat.Mensaje, servidor *claseschat.Servidor) {
	var aux net.Conn
	for i := 0; i < len(servidor.Usuarios); i++ {
		if servidor.Usuarios[i].Nombre != msg.Enviador {
			aux = servidor.Usuarios[i].Conexion
			servidor.Usuarios[i].MensajesRecibidos = append(servidor.Usuarios[i].MensajesRecibidos, *msg)
			err := gob.NewEncoder(aux).Encode(msg)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

}


func main() {
	var tematica string
	lservidor := new(claseschat.Servidor)
	fmt.Println("Tematica del servidor: ")
	fmt.Scan(&tematica)
	lservidor.Tematica = tematica
	go servidor(lservidor) //inicio de serviodr
	var op int
	Detener := false
	for !Detener {
		fmt.Println("Opciones de servidor: ")
		fmt.Println("1-Mostrar todos los mensajes ")
		fmt.Println("3-Terminar servidor ")
		fmt.Scan(&op)
		if op == 1 {
			lservidor.MostrarTodosLosMensajes()
		}  else if op == 3 {
			Detener = true
		}
	}

}
