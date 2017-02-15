package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	mundo := "mundo"
	edad := 20
	nombre := "Jonathan"
	bandera := true
	peso := 60.45

	fmt.Println("Hola ", mundo) // <---- Forma más fácil
	//Formas con printf
	fmt.Printf("Mi nombre es: %v \n", nombre) // %s
	fmt.Printf("Mi edad es: %d \n", edad)
	fmt.Printf("¿tengo sueño? --> %t \n", bandera)
	fmt.Printf("Mi peso es: %f \n", peso) // %e   %b  --> son para científicos

	//Leer datos forma 1
	var edad_ingresa int
	fmt.Print("Ingresa de nuevo tu edad: ")
	fmt.Scanln(&edad_ingresa)
	fmt.Println("edad ingresada: ", edad_ingresa)

	//Leer datos con scanf forma 2
	var edad_f int
	fmt.Print("Coloca tu edad f: ")
	fmt.Scanf("%d\n", &edad_f)
	fmt.Printf("La edad f es: %d\n", edad_f)

	//Forma 3 con bufio
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingresa tu nombre: ")
	texto, err := reader.ReadString('\n') // ---> lee hasta que haya un salto de línea
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Hola ", texto)
	}

}
