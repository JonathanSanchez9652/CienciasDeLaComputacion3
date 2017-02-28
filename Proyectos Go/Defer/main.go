// Defer project main.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	ejecucion := readFile()
	fmt.Println(ejecucion)
}

func readFile() bool {
	archivo, error := os.Open("./ArchivoGo.txt")
	defer func() {
		archivo.Close()
		fmt.Println("Defer")
	}()

	if error != nil {
		fmt.Println("Hubo un error")
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
