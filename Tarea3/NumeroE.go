package main

import "fmt"

func main() {
	var n int
	fmt.Println("Ingresa el numero de veces que quieres que se repita: ")
	fmt.Scanf("%d", &n)
	var resultado float32 = 1
	for i := 1; i <= n; i++ {
		resultado = resultado + (1 / float32(factorial(i)))
	}
	fmt.Printf("resultado: %f", resultado)
}

func factorial(numero int) int {
	factorial := numero
	for i := numero - 1; i > 0; i-- {
		factorial = factorial * i

	}
	return factorial
}
