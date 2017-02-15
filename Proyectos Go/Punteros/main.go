// Punteros project main.go
package main

import (
	"fmt"
)

func main() {
	/*
		1.Es una dirección de memoria
		2.En lugar del valor, tenemos la dirección donde está el valor
		3.X,Y => as1321j => 5
		4.X => 6
		5.Y ¿? => 6
		*T => *int, *string, *struct
	*/
	var x, y *int
	entero := 5
	x = &entero //--> retorna la dirección donde está 5
	y = &entero

	fmt.Println("Dirección de memoria de x: ", x)
	fmt.Println("Dirección de memoria de y: ", y)
	fmt.Println("Valor de x: ", *x)
	fmt.Println("Valor de y: ", *y)

	//Modificación de datos
	*x = 6
	//La posición sigue siendo la misma, pero el valor si cambia
	fmt.Println("Modificación de datos")
	fmt.Println("Dirección de memoria de x: ", x)
	fmt.Println("Dirección de memoria de y: ", y)
	fmt.Println("Valor de x: ", *x)
	fmt.Println("Valor de y: ", *y)

}
