// MakeYAppend project main.go
package main

import (
	"fmt"
)

func main() {
	//nombre slice := make(tipo,tamaño,capacidad-->OPCIONAL)
	slice := make([]int, 3, 5)
	fmt.Println("Contenido del slice: ", slice)
	fmt.Println("Tamaño del slice: ", len(slice))
	fmt.Println("Capacidad / CAPACITY / del slice: ", cap(slice))

	//agregando un nuevo elemento
	slice = append(slice, 2) //--> agrega el elemento 2
	fmt.Println("Slice con el nuevo elementro agregado: ", slice)

}
