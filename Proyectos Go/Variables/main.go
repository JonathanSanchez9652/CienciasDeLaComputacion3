package main

import "fmt"

func main() {
	//Todas las variables deben ser utilizadas

	//Forma 1 de declarar variables
	var x int
	x = 23

	//Forma 2 de declarar variables
	y := 24
	nombre := "coco"

	fmt.Println("Numero x =", x)
	fmt.Println("Nombre y =", y)
	fmt.Println("Nombre cadena =", nombre)

	//Para re asignar variables es necesario usar "="
	modificable := "inicial"
	fmt.Println(modificable)

	modificable = "nueva"
	fmt.Println(modificable)

}
