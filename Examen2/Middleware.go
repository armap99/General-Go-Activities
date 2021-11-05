package main

import (
	"fmt"
	"net"
	"net/rpc"
	"strconv"

	"./claseschat"
)

type Server struct{
	PuertoLibre string
	Servidores []claseschat.Servidor
}

func (this *Server) ObtenerPuertoLibre (data int ,reply *string) error{
	*reply = this.PuertoLibre
	intStr, _ := strconv.Atoi(this.PuertoLibre,)
	intStr += 1
	this.PuertoLibre = strconv.Itoa(intStr)
	return nil

}

func (this *Server) AgregarSala (data claseschat.Servidor, reply *string) error{
	this.Servidores = append(this.Servidores, data)
	*reply = "Servidor agregado"
	this.ImprimirServidoresL()
	return nil
}

func (this *Server) ImprimirServidoresL () {
	for i := 0; i < len(this.Servidores); i++ {
		fmt.Println("Servidor ", i, " Tematica: ", this.Servidores[i].Tematica, "127.0.0.1 Puerto: ",this.Servidores[i].Puerto, " Participantes: ", this.Servidores[i].Participantes)
	}

}

func (this *Server) ImprimirServidores (data int,reply *string) error {
	for i := 0; i < len(this.Servidores); i++ {
		fmt.Println("Servidor ", i, " Tematica: ", this.Servidores[i].Tematica, "127.0.0.1 Puerto: ",this.Servidores[i].Puerto, " Participantes: ", this.Servidores[i].Participantes)
	}
	*reply = ""
	return nil
}

func (this *Server) RegresarServidores (data int,reply *string) error{
	aux := ""
	for i := 0; i < len(this.Servidores); i++ {
		aux += "Servidor "+ strconv.Itoa(i) + " Tematica: " + this.Servidores[i].Tematica + "127.0.0.1 Puerto: " + this.Servidores[i].Puerto + " Participantes: " + strconv.Itoa(this.Servidores[i].Participantes)
	}
	*reply = aux

	return nil
}

func (this *Server) ObtenerPuertoPorTema (data string,reply *string) error{
	aux := ""
	for i := 0; i < len(this.Servidores); i++ {
		if this.Servidores[i].Tematica == data {
			aux = this.Servidores[i].Puerto
			this.Servidores[i].Participantes += 1
			break
		}
		
	}
	*reply = aux

	return nil
}

/////////////////////////////////////////////////////////////////////////////////////////////
//main

func server() {
	servidor := new(Server)
	servidor.PuertoLibre = "8000"
	rpc.Register(servidor)
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Corriendo Middleware...")
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(c)
	}
}

func main(){
	go server()

	var input string
	fmt.Scanln(&input)
}
