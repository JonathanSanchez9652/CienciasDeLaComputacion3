package main

import "fmt"

func main() {
	//Forma 1 de declararlo
	var arreglo [10]int
	//imprime todo el arreglo
	// si no se declaran valores iniciales el valor por defecto es
	//0 para int, '' para string y false para bool
	fmt.Println(arreglo)

	//forma 2 de declararlo con valores iniciales
	arreglo2 := [3]int{1, 2, 4}
	fmt.Println(arreglo2) //--> imprime todo el arreglo
	//tamaño del arreglo
	fmt.Println("El tamaño del arreglo es: ", len(arreglo2))

	//iterar el arreglo
	for i := 0; i < len(arreglo2); i++ {
		fmt.Println("dato ", arreglo2[i], " en la posición ", i)
	}

	//asignar un nuevo valor al arreglo
	arreglo2[1] = 40

	//volver a imprimir
	fmt.Println("Arreglo modificado: ", arreglo2)

	//ARREGLO MULTI DIMENSIONAL
	var matriz [2][2]int
	fmt.Println("matriz vacía: ", matriz)

}
