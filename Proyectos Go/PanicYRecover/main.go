// Defer project main.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	executeReadFile()
	fmt.Println("Nunca me voy a imprimir")
}

func executeReadFile() {
	ejecucion := readFile()
	fmt.Println(ejecucion)
}

func readFile() bool {
	archivo, error := os.Open("./ArchivooGo.txt")
	defer func() {
		archivo.Close()
		fmt.Println("Defer")

		r := recover() // evita el panic
		fmt.Println(r)
	}()

	if error != nil {
		panic(error) //forma en la cual se puede imprimir el error
	}
	var i int
	fmt.Println("Línea | Texto")
	fmt.Println("--------------------")
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		i++
		linea := scanner.Text()
		fmt.Println(i, "     |", linea)
	}

	return true
}
