package main

import "fmt"

func main() {
	var base int
	var altura int
	fmt.Println("Ingresa el valor de la base del trinagulo: ")
	fmt.Scanf("%d", &base) //fmt.Scan(&lado)
	fmt.Println("Ingresa el valor de la altura del trinagulo: ")
	fmt.Scan(&altura)
	resultado := (base * altura) / 2
	fmt.Println("El area del triandulo es: ", resultado)
}
