package main

import "fmt"

// interfas animal con el metodo move
type animal interface {
	move() string
}

type dog struct{}

type cat struct{}

// funcion que pertenece a la estructura (dog)
// e interfaz animal que contiene el metodo move()
// por lo que implicitamente dog pertenece a la interfaz animal
// retorna el mensaje a imprimir
func (dog) move() string {
	return ("I'm a dog and I'm swimming")
}

// funcion que pertenece a la estructura (cat)
// e interfaz animal que contiene el metodo move()
// por lo que implicitamente cat pertenece a la interfaz animal
// retorna el mensaje a imprimir
func (cat) move() string {
	return ("I'm a cat and I'm flying")
}

// actionAnimal: funcion que imprime el mensaje del animal
// a: de tipo animal
func actionAnimal(a animal) {
	fmt.Println(a.move())
}
func main() {
	c := cat{} // instanciuna structura cat{} y la asigno a c
	p := dog{}

	actionAnimal(c) // ejecuta la funcion actionAnimal en c
	actionAnimal(p)
}
