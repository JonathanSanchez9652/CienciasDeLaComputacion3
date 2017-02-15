// Structs project main.go
package main

import (
	"fmt"
)

type User struct {
	edad     int
	nombre   string
	apellido string
}

func main() {
	jonathan := User{edad: 20, nombre: "Jonathan", apellido: "Sánchez"}
	javier := User{14, "R", "f"}

	fmt.Println(jonathan)
	fmt.Println(jonathan.nombre)
	fmt.Println(javier)
	fmt.Println(javier.apellido)

	jonathan.edad = 25
	fmt.Println(jonathan.edad)
}
