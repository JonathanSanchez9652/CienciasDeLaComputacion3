package main

import (
	"fmt"
	"regexp"

	"strings"
)

func main() {
	//var entrada string
	//fmt.Scanln(&entrada)
	entrada := "5 3 >= false 8 3 = &"
	InsertarPila(entrada)
}

func InsertarPila(a string) {
	val := strings.Split(a, " ")
	arreglo := make([]string, len(val))
	for i := 0; i < len(val); i++ {
		var r1 = regexp.MustCompile(`^[0-9]+$`)
		t1 := r1.MatchString(val[i])
		if t1 == true {
			arreglo[i] = "VN"
			fmt.Println(arreglo[i], " ", val[i])
		}

		var r2 = regexp.MustCompile("false")
		t2 := r2.FindString(val[i])
		if t2 == "false" || t2 == "true" {
			arreglo[i] = "VL"
			fmt.Println(arreglo[i], " ", val[i])
		}

		var r2a = regexp.MustCompile("true")
		t2a := r2a.FindString(val[i])
		if t2a == "false" || t2a == "true" {
			arreglo[i] = "VL"
			fmt.Println(arreglo[i], " ", val[i])
		}
		var r3 = regexp.MustCompile("=|<|>")
		t3 := r3.MatchString(val[i])
		if t3 == true {
			arreglo[i] = "OC"
			fmt.Println(arreglo[i], " ", val[i])
		}

		var r4 = regexp.MustCompile("&")
		t4 := r4.MatchString(val[i])
		if t4 == true {
			arreglo[i] = "OL"
			fmt.Println(arreglo[i], " ", val[i])
		}

	}

}
