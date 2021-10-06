package main

import "fmt"

func main() {
	var radio float32
	fmt.Println("Ingresa el valor del radio del circulo: ")
	fmt.Scanf("%f", &radio) //fmt.Scan(&lado)
	resultado := 3.1416 * (radio * radio)
	fmt.Println("El area del circulo es: ", resultado)
}
