package main

import "fmt"

func main() {
	//forma 1 con todos los parámetros
	for i := 0; i < 10; i++ { //los 3 son opcionales
		fmt.Println("Hola mundo", i)
	}

	//simulando un "while"
	j := 0
	for j < 10 {
		fmt.Println("hola pseudo while", j)
		j++
	}

	//ciclo "infinito" ---> uso del break
	h := 0
	var terminar string

	for {
		fmt.Println("hola 'infinito'", h)
		h++

		if h > 7 {
			fmt.Println("Finalizar S/N?")
			fmt.Scan(&terminar)
			fmt.Println(terminar)
			if terminar == "S" {
				break
			} else {
				h = 0
			}

		}
	}

	//ciclo infinito ---> uso del continue
	l := 0

	for {
		if l == 2 {
			l++
			continue //continua volviendo al inicio del ciclo
		}
		fmt.Println("continue", l)
		l++
		if l > 10 {
			break
		}
	}

}
