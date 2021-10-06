package main

import "fmt"

func main() {
	var Dia int
	var Mes int
	fmt.Println("Ingresa tu dia de nacimiento: ")
	fmt.Scan(&Dia)
	fmt.Println("Ingresa tu mes de necimiento: ")
	fmt.Scan(&Mes)
	switch Mes {
	case 1:
		if Dia <= 20 {
			fmt.Printf("Eres Capricornio")
		} else if Dia >= 21 {
			fmt.Printf("Eres Acuario")
		}
	case 2:
		if Dia <= 18 {
			fmt.Printf("Eres Acuario")
		} else if Dia >= 19 {
			fmt.Printf("Eres Piscis ")
		}
	case 3:
		if Dia <= 20 {
			fmt.Printf("Eres Piscis")
		} else if Dia >= 21 {
			fmt.Printf("Eres Aries")
		}
	case 4:
		if Dia <= 20 {
			fmt.Printf("Eres Aries")
		} else if Dia >= 21 {
			fmt.Printf("Eres Tauro")
		}
	case 5:
		if Dia <= 21 {
			fmt.Printf("Eres Tauro")
		} else if Dia >= 22 {
			fmt.Printf("Eres Géminis")
		}
	case 6:
		if Dia <= 21 {
			fmt.Printf("Eres Géminis")
		} else if Dia >= 22 {
			fmt.Printf("Eres Cáncer")
		}
	case 7:
		if Dia <= 22 {
			fmt.Printf("Eres Cáncer")
		} else if Dia >= 23 {
			fmt.Printf("Eres Leo")
		}
	case 8:
		if Dia <= 23 {
			fmt.Printf("Eres Leo")
		} else if Dia >= 24 {
			fmt.Printf("Eres Virgo")
		}
	case 9:
		if Dia <= 23 {
			fmt.Printf("Eres Virgo ")
		} else if Dia >= 24 {
			fmt.Printf("Eres Libra ")
		}
	case 10:
		if Dia <= 23 {
			fmt.Printf("Eres Libra ")
		} else if Dia >= 24 {
			fmt.Printf("Eres Escorpión ")
		}
	case 11:
		if Dia <= 22 {
			fmt.Printf("Eres Escorpión ")
		} else if Dia >= 23 {
			fmt.Printf("Eres Sagitario")
		}
	case 12:
		if Dia <= 21 {
			fmt.Printf("Eres Sagitario")
		} else if Dia >= 22 {
			fmt.Printf("Eres Capricornio")
		}
	}
}
