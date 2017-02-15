// GoRoutines project main.go
package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	go miNombreLento("Jonathan") //el go lo uso para que el programa sea síncrono
	fmt.Println("Finalizoooooooooooooo")
	//Se usa para que ejecute todo y lo termine a menos que yo oprima enter
	var wait string
	fmt.Scanln(&wait)
}
func miNombreLento(name string) {
	letras := strings.Split(name, "")
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}
