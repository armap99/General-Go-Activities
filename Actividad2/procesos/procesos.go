package procesos

import (
	"fmt"
	"time"
)

type Bloque struct {
	Procesos []Proceso
}

func (s *Bloque) Start() {
	go func() {
		for {
			for i := 0; i < len(s.Procesos); i++ {
				s.Procesos[i].Tiempo += 1
			}
			time.Sleep(time.Millisecond * 500)
		}
	}()
	var input string
	fmt.Scanln(&input)
}

func (s *Bloque) MostrarGeneral(c chan string) {
	go func() {
		for {
			select {
			case msg := <-c:
				if msg == "detener" {
					return
				}
			default:
				for i := 0; i < len(s.Procesos); i++ {

					s.Procesos[i].Mostrar()
				}
				time.Sleep(time.Millisecond * 500)
			}

		}
	}()
	var input string
	fmt.Scanln(&input)
}

func (s *Bloque) Stop(c chan string) {
	c <- "detener"
}

func (s *Bloque) Eliminar(id uint64) []Proceso {
	for i := 0; i < len(s.Procesos); i++ {
		if s.Procesos[i].Id == id {
			s.Procesos = s.RemoveIndex(i)
			return s.Procesos
		}
	}
	return s.Procesos
}

func (s *Bloque) RemoveIndex(index int) []Proceso {
	return append(s.Procesos[:index], s.Procesos[index+1:]...)
}

type Proceso struct {
	Id      uint64
	Tiempo  uint64
	Estatus int
}

func (p *Proceso) Mostrar() {
	fmt.Println(p.Id, ": ", p.Tiempo)

}
