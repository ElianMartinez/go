package main

import (
	"fmt"
)

type Persona struct {
	Nombre string
	Edad   int
}

func (p Persona) Saludar() {
	fmt.Println("Hola soy ", p.Nombre)
}

func main() {
	var persona Persona = Persona{Nombre: "Juan", Edad: 25}

	fmt.Println(persona.Nombre)
	fmt.Println(persona.Edad)

}
