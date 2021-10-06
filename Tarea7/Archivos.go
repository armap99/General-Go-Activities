package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	var n uint
	var s []string
	var aux string
	var op int
	fmt.Println("Ingresa la cantidad de numeros de tu slice: ")
	fmt.Scan(&n)
	for i := 0; i < int(n); i++ {
		fmt.Println("Ingresa un string: ")
		fmt.Scan(&aux)
		s = append(s, aux)
	}
	fmt.Println("1-Asecendente\n2-Descendente")
	fmt.Scan(&op)
	if op == 1 {
		file, err := os.Create("asecendente.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		sort.Strings(s)
		for _, v := range s {
			file.WriteString(v)
			file.WriteString("\n")
		}
	} else if op == 2 {
		file, err := os.Create("descendente.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		sort.Sort(sort.Reverse(sort.StringSlice(s)))
		for _, v := range s {
			file.WriteString(v)
			file.WriteString("\n")
		}

	}

}
