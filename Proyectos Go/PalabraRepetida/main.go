// PalabraRepetida project main.go
package main

import (
	"fmt"
)

func main() {

	//nombre slice := make(tipo,tamaño,capacidad-->OPCIONAL)
	slice := make([]string, 0, 10)

	slice = append(slice, "a")
	slice = append(slice, "b")
	slice = append(slice, "c")
	slice = append(slice, "d")
	slice = append(slice, "e")
	slice = append(slice, "f")
	slice = append(slice, "f")
	slice = append(slice, "a")

	var palabra string
	var contador int
	var estaRepetida bool
	sliceSinRepetidas := make([]string, 1, 20)

	for i := 0; i < len(slice); i++ {
		palabra = slice[i]

		for k := 0; k < len(sliceSinRepetidas); k++ {
			if sliceSinRepetidas[k] == palabra {
				estaRepetida = true
				break
			}
			estaRepetida = false
		}

		if estaRepetida == false {
			for j := 0; j < len(slice); j++ {
				if slice[j] == palabra {
					contador++
				}
			}
			sliceSinRepetidas = append(sliceSinRepetidas, palabra)
		} else {
			i++
			continue
		}

		fmt.Println(palabra, "->", contador, "veces")
		contador = 0

		estaRepetida = false

	}

	fmt.Println("Slice: ", slice)
	fmt.Println("Slice sin repetidas: ", sliceSinRepetidas)
	fmt.Println("Palabra: ", palabra)
	fmt.Println("¿está repetida?: ", estaRepetida)

}
