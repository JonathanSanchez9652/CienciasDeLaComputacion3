// Copy project main.go
package main

import (
	"fmt"
)

func main() {

	//Copy funciona para copiar dos arreglos
	//copy(destino,fuente)
	//Copia el mínimo de elementos en ambos arreglos
	slice := []int{1, 2, 3, 4}
	//--> Es la mejor manera para siempre copiar el mismo tamaño del original len()
	//una práctica común es duplicar la capacidad
	copia := make([]int, len(slice), cap(slice)*2)

	fmt.Println("Original: ", slice)
	fmt.Println("Copia: ", copia)

	copy(copia, slice)
	fmt.Println("Hecha la copia")
	fmt.Println("Original: ", slice)
	fmt.Println("Copia: ", copia)

}
