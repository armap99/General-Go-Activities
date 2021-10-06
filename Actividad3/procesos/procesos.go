package procesos

import (
	"fmt"
	"time"
)

type Bloque struct {
	Procesos []Proceso
}

func (s *Bloque) Primero() Proceso {
	return s.Procesos[0]
}

func (s *Bloque) Start() {
	go func() {
		for {
			for i := 0; i < len(s.Procesos); i++ {
				s.Procesos[i].Tiempo += 1
				s.Procesos[i].Mostrar()
			}
			fmt.Println("-------------------------", len(s.Procesos))
			time.Sleep(time.Millisecond * 500)
		}
	}()
}

func (s *Bloque) RemoveIndex(index int) {
	fmt.Println("entrada")
	s.Procesos = append(s.Procesos[:index], s.Procesos[index+1:]...)
	fmt.Println("salida")
}

type Proceso struct {
	Id      uint64
	Tiempo  uint64
	Estatus int
}

func (p *Proceso) Mostrar() {
	fmt.Println(p.Id, ": ", p.Tiempo)

}
