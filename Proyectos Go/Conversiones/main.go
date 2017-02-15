package main

//todos los paquetes nombrados deben ser utilizados
import (
	"fmt"
	"strconv"
)

func main() {
	//convertir de entero a string
	edad := 22
	edad_str := strconv.Itoa(edad)
	fmt.Println("Mi edad es: " + edad_str)

	//convertir de string a entero
	anios := "23"
	anios_int, _ := strconv.Atoi(anios)
	fmt.Println("Mi edad 2 es: ", anios_int)

}
