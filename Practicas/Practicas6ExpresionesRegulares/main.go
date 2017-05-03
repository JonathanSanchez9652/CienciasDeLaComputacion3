// Practicas6ExpresionesRegulares project main.go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	var texto string
	fmt.Print("Ingresa el dato: ")
	fmt.Scanln(&texto)
	
	var validNum = regexp.MustCompile(`^[0-9]+$`)
	veriNum := validNum.MatchString(texto)
	
	validLogico = 
	
	if veriNum == true {
		fmt.Println("UN -> ", texto)
	}

}
