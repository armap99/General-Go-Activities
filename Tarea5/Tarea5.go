package main

import "fmt"

func Ordenamiento(arreglo []int) {
	var aux int
	for j := 0; j < len(arreglo); j++ {
		for i := 0; i < len(arreglo)-1; i++ {
			if arreglo[i] > arreglo[i+1] {
				aux = arreglo[i]
				arreglo[i] = arreglo[i+1]
				arreglo[i+1] = aux
			}
		}
	}

	fmt.Println("El slice odenado es: ", arreglo)
}

func fibonacci(i int) int {
	if i <= 1 {
		return i
	}

	return fibonacci(i-1) + fibonacci(i-2)
}

func generadorInpares() func() uint {
	i := uint(1)
	return func() uint {
		var inpar = i
		i += 2
		return inpar
	}
}

func intercambia(a *int, b *int) {
	var aux int
	aux = *a
	*a = *b
	*b = aux
}

func main() {
	var op int
	fmt.Println("1 - Ordenamiento de slice \n 2 - Secuencia de Fibonacci \n 3 - Generador de numeros impares \n 4 - Intercambio de valor \n Igrese la opcion: ")
	fmt.Scan(&op)
	if op == 1 {
		var n uint
		var s []int
		var aux int
		fmt.Println("Ingresa la cantidad de numeros de tu slice: ")
		fmt.Scan(&n)
		for i := 0; i < int(n); i++ {
			fmt.Println("Ingresa un numero: ")
			fmt.Scan(&aux)
			s = append(s, aux)
		}
		Ordenamiento(s)
	} else if op == 2 {
		var x int
		fmt.Println("Ingrese el largo de la serie: ")
		fmt.Scan(&x)
		fmt.Println("\nSerie: ")
		for i := x; i != 0; i-- {
			fmt.Println(fibonacci(i))
		}

	} else if op == 3 {
		var n int
		nextPar := generadorInpares()
		fmt.Println("Ingresa la cantidad de numeros impares que quieres: ")
		fmt.Scan(&n)
		for i := 0; i < int(n); i++ {
			fmt.Println(nextPar())
		}
	} else if op == 4 {
		var numero1 int
		var numero2 int
		fmt.Println("Ingresa el primer numero que quieres cambiar : ")
		fmt.Scan(&numero1)
		fmt.Println("Ingresa el segundo numero que quieres cambiar : ")
		fmt.Scan(&numero2)
		intercambia(&numero1, &numero2)
		fmt.Println("A: ", numero1)
		fmt.Println("B: ", numero2)

	}

}
