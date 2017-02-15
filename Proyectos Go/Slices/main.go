// Slices project main.go
package main

import (
	"fmt"
)

func main() {
	//Los slices a diferencia de los arreglos no tiene tamaño fijo
	//Definición de slices
	//Los slices sin inicializar sus valor por defecto es nulo
	matriz := []int{1, 2, 3, 4}
	fmt.Println(matriz)

	var matrizVacia []int
	if matrizVacia == nil {
		fmt.Println("La matriz está vacía", matrizVacia)
	} else {
		fmt.Println(matrizVacia)
	}
	//longitud del slice
	fmt.Println("Longitud del slice: ", len(matriz))

	//Slices a partir de un arreglo
	arreglo := [5]int{1, 2, 3, 4, 5}
	slice := arreglo[:2]   //del inicio a la posición 2
	slice2 := arreglo[2:4] // de la posición 2 a la 4
	fmt.Println("Slice1: ", slice)
	fmt.Println("Slice2: ", slice2)

}
