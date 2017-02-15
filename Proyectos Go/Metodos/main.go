// Metodos project main.go
package main

import (
	"fmt"
)

type User struct {
	edad     int
	nombre   string
	apellido string
}

//función(nombreVarQueRecibe Tipo) nombreFuncion() tipoRetorno{
func (usuario User) nombre_completo() string {
	return usuario.nombre + " " + usuario.apellido
}

//Método para cambiar el nombre
func (this *User) set_name(n string) {
	this.nombre = n
}
func main() {
	jonathan := new(User)
	jonathan.edad = 25
	jonathan.nombre = "Jonathan"
	jonathan.apellido = "Sánchez"

	jonathan.set_name("Javier") //--> instrucción y llamado al método para cambiar nombre

	fmt.Println(jonathan.nombre_completo())
	fmt.Println(jonathan.nombre)

}
