package main

import (
	"fmt"
)

func main() {
	/* OPERADORES
	== igual a
	< menor que			<= menor ó igual que
	> mayor que			>= mayor ó igual que
	&&  AND
	||  OR
	*/
	x := 5
	y := 5
	if x > y { //tienen que ir en el mismo renglón ...{
		fmt.Println(x, "es mayor que", y)
	} else if x < y {
		fmt.Println(x, "es menor que", y)
	} else {
		fmt.Println(x, "y", y, "son iguales")
	}

}
