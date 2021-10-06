package main

import (
	"fmt"
	"sort"
)

type Proceso struct {
	id        uint64
	Prioridad int64
	Tiempo    uint64
	Estatus   string
}

type ByPrioridad []Proceso

func (a ByPrioridad) Len() int           { return len(a) }
func (a ByPrioridad) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPrioridad) Less(i, j int) bool { return a[i].Prioridad < a[j].Prioridad }

type ByPrioridadDes []Proceso

func (a ByPrioridadDes) Len() int           { return len(a) }
func (a ByPrioridadDes) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPrioridadDes) Less(i, j int) bool { return a[i].Prioridad > a[j].Prioridad }

func (i *Proceso) Mostrar() {
	fmt.Println("Id: ", i.id)
	fmt.Println("Prioridad: ", i.Prioridad)
	fmt.Println("Tiempo:", i.Tiempo)
	fmt.Println("Estatus: ", i.Estatus)
	fmt.Println(" ")
}

func main() {
	s := []Proceso{
		{id: 5, Prioridad: 8, Tiempo: 15, Estatus: "Aceptado"},
		{id: 1, Prioridad: 5, Tiempo: 30, Estatus: "Espera"},
		{id: 4, Prioridad: 1, Tiempo: 5, Estatus: "Bloqueado"},
		{id: 3, Prioridad: 6, Tiempo: 15, Estatus: "Aceptado"},
		{id: 2, Prioridad: 10, Tiempo: 27, Estatus: "Espera"},
	}
	fmt.Println("Sin ordenar: ")
	for _, v := range s {
		v.Mostrar()
	}
	fmt.Println("Asecendente: ")
	sort.Sort(ByPrioridad(s))
	for _, v := range s {
		v.Mostrar()
	}
	fmt.Println("Descendente: ")
	sort.Sort(ByPrioridadDes(s))
	for _, v := range s {
		v.Mostrar()
	}
}
