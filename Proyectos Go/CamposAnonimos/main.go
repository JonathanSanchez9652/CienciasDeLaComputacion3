// CamposAnonimos project main.go
package main

import (
	"fmt"
)

type Human struct {
	name string
}

type Dummy struct {
	name string
}
type Tutor struct {
	Human
	Dummy
}

func (this Human) hablar() string {
	return "bla bla bla"
}
func (this Tutor) hablaTutor() string {
	return this.Human.hablar() + " Bienvenido"
}
func main() {
	tutor := Tutor{Human{"Jonathan"}, Dummy{"hhhh"}}
	fmt.Println(tutor.Human.name)
	fmt.Println(tutor.Dummy.name)
	fmt.Println(tutor.hablaTutor())

}
