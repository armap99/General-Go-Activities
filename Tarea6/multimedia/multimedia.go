package multimedia

import "fmt"

type ContenidoWeb struct {
	Multimedias []Multimedia
}

func (cw *ContenidoWeb) Mostrar() {
	for _, f := range cw.Multimedias {
		f.Mostrar()
	}
}

type Multimedia interface {
	Mostrar()
}

type Imagen struct {
	Titulo  string
	Formato string
	Canal   string
}

func (i *Imagen) Mostrar() {
	fmt.Println("Imagen")
	fmt.Println("Titulo: ", i.Titulo)
	fmt.Println("Formato: ", i.Formato)
	fmt.Println("Canal: ", i.Canal)
	fmt.Println(" ")
}

type Audio struct {
	Titulo   string
	Formato  string
	Duracion int
}

func (a *Audio) Mostrar() {
	fmt.Println("Audio")
	fmt.Println("Titulo: ", a.Titulo)
	fmt.Println("Formato: ", a.Formato)
	fmt.Println("Duracion: ", a.Duracion)
	fmt.Println(" ")
}

type Video struct {
	Titulo  string
	Formato string
	Frames  int
}

func (v *Video) Mostrar() {
	fmt.Println("Video")
	fmt.Println("Titulo: ", v.Titulo)
	fmt.Println("Formato: ", v.Formato)
	fmt.Println("Frames: ", v.Frames)
	fmt.Println(" ")
}
