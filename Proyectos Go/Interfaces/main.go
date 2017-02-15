// Interfaces project main.go
package main

import (
	"fmt"
)

// INTERFAZ DE UN USUARIO EN GENERAL
type User interface {
	Permisos() int // 1-5
	Nombre() string
}

// ESTRUCTURA PARA ADMIN
type Admin struct {
	nombre string
}

func (this Admin) Permisos() int {
	return 5
}

func (this Admin) Nombre() string {
	return this.nombre
}

//ESTRUCTURA PARA EDITOR
type Editor struct {
	nombre string
}

func (this Editor) Permisos() int {
	return 3
}
func (this Editor) Nombre() string {
	return this.nombre
}

// AUTENTICA EL TIPO DE USUARIO
func autenticacion(user User) string {
	permisos := user.Permisos()
	if permisos > 4 {
		return user.Nombre() + " tiene permisos de administrador"
	} else if permisos == 3 {
		return user.Nombre() + " tiene permisos de editor"
	}
	return ""
}

func main() {
	admin := Admin{"Jonathan"} // Instancia de un admin
	editor := Editor{"Javier"} //Instancia de un editor
	fmt.Println(autenticacion(admin))
	fmt.Println(autenticacion(editor))

	//otra forma almacenando las estructuras en un arreglo
	usuarios := []User{Admin{"Ruth"}, Editor{"López"}}
	for _, usuario := range usuarios {
		fmt.Println(autenticacion(usuario))
	}
}
