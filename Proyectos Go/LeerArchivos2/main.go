// LeerArchivos2 project main.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//lee el archivo línea por línea
	archivo, error := os.Open("./ArchivoGo.txt")

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
	archivo.Close()
}
