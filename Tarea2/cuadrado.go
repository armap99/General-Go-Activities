package main

import "fmt"

func main() {
	var lado int
	fmt.Println("Ingresa el valor de los lados del cuadrado: ")
	fmt.Scanf("%d", &lado) //fmt.Scan(&lado)
	resultado := lado * lado
	fmt.Println("El area del cuadrado es: ", resultado)
}
