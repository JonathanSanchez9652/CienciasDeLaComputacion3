// LeerArchivos1 project main.go
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//lee el archivo cargado en memoria
	fileData, err := ioutil.ReadFile("./ArchivoGo.txt")

	if err != nil {
		fmt.Println("hubo un error")
	}
	fmt.Println(string(fileData))
}
