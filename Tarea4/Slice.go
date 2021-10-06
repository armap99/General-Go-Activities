package main

import "fmt"

func main() {
	var n uint
	var s []int
	var aux int
	var total int
	fmt.Println("Ingresa la cantidad de numeros que vas a sumar: ")
	fmt.Scan(&n)
	for i := 0; i < int(n); i++ {
		fmt.Println("Ingresa un numero: ")
		fmt.Scan(&aux)
		s = append(s, aux)
	}
	for _, v := range s {
		total = total + v
	}
	println("Total: ", total)

}
