package main

import "fmt"

func main() {
	var f float32
	fmt.Println("Ingresa el valor de los grados Fahrenheit que quiere convertir a Celsius: ")
	fmt.Scanf("%f", &f) //fmt.Scan(&lado)
	var resultado float32 = (f - 32) * 5 / 9
	fmt.Println("Celsius: ", resultado)
}
