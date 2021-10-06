package main

import (
	"fmt"

	"./multimedia"
)

func main() {
	var aux int = 100
	contenido := new(multimedia.ContenidoWeb)
	for i := 0; i < aux; i++ {
		var op int
		fmt.Println("Bienvenido")
		fmt.Println("1-Agregar una imagen")
		fmt.Println("2-Agregar un audio")
		fmt.Println("3-Agregar un video")
		fmt.Println("4-Mostrar")
		fmt.Println("5-Salir")
		fmt.Scan(&op)
		if op == 1 {
			var titulo string
			var formato string
			var canal string
			fmt.Println("Ingresa el titulo de la imagen: ")
			fmt.Scan(&titulo)
			fmt.Println("Ingresa el formato de la imagen: ")
			fmt.Scan(&formato)
			fmt.Println("Ingresa el canal de la imagen: ")
			fmt.Scan(&canal)
			image := multimedia.Imagen{Titulo: titulo, Formato: formato, Canal: canal}
			contenido.Multimedias = append(contenido.Multimedias, &image)
		} else if op == 2 {
			var titulo string
			var formato string
			var duracion int
			fmt.Println("Ingresa el titulo del audio: ")
			fmt.Scan(&titulo)
			fmt.Println("Ingresa el formato del audio: ")
			fmt.Scan(&formato)
			fmt.Println("Ingresa la duracion del audio: ")
			fmt.Scan(&duracion)
			audi := multimedia.Audio{Titulo: titulo, Formato: formato, Duracion: duracion}
			contenido.Multimedias = append(contenido.Multimedias, &audi)
		} else if op == 3 {
			var titulo string
			var formato string
			var frames int
			fmt.Println("Ingresa el titulo del video: ")
			fmt.Scan(&titulo)
			fmt.Println("Ingresa el formato del video: ")
			fmt.Scan(&formato)
			fmt.Println("Ingresa los frames del video: ")
			fmt.Scan(&frames)
			vid := multimedia.Video{Titulo: titulo, Formato: formato, Frames: frames}
			contenido.Multimedias = append(contenido.Multimedias, &vid)
		} else if op == 4 {
			fmt.Println("Lista: ")
			contenido.Mostrar()
		} else if op == 5 {
			aux = 101
		}

	}
}
