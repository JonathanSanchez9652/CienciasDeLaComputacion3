// Practica3 project main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Arbol struct {
	Izq   *Arbol
	Valor string
	Der   *Arbol
}

func NewArbol() *Arbol {
	return &Arbol{}
}

type Stack struct {
	arb   []*Arbol
	count int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(n *Arbol) {
	s.arb = append(s.arb[:s.count], n)
	s.count++
}

func (s *Stack) Pop() *Arbol {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.arb[s.count]
}

func Operacion(a *Arbol) int {
	if a == nil {
		return 0
	} else if a.Valor == "+" {
		return Operacion(a.Izq) + Operacion(a.Der)
	} else if a.Valor == "-" {
		return Operacion(a.Izq) - Operacion(a.Der)
	} else if a.Valor == "*" {
		return Operacion(a.Izq) * Operacion(a.Der)
	} else if a.Valor == "/" {
		if Operacion(a.Der) == 0 {
			fmt.Println("***************************")
			fmt.Println("ERROR DIVISIÓN POR CERO")
			fmt.Println("***************************")
		} else {
			return Operacion(a.Izq) / Operacion(a.Der)
		}
		return 0
	} else {
		conv, err := strconv.Atoi(a.Valor)
		if err != nil {
			fmt.Println("")
			fmt.Println(".....................")
			fmt.Println("Error de sintaxis ", a.Valor, "no es un caracter válido")
		}
		return conv
	}

}

func InOrden(a *Arbol) {
	if a == nil {
		return
	}
	InOrden(a.Izq)
	fmt.Printf(a.Valor)
	InOrden(a.Der)
}

func InsertarPila(a string) *Arbol {
	stack := NewStack()
	val := strings.Split(a, " ")
	for i := 0; i < len(val); i++ {
		arbol := &Arbol{nil, val[i], nil}
		if arbol.Valor != "+" && arbol.Valor != "-" && arbol.Valor != "*" && arbol.Valor != "/" {
			stack.Push(arbol)
		} else {
			arbol.Der = stack.Pop()
			arbol.Izq = stack.Pop()
			stack.Push(arbol)
		}
	}
	return stack.Pop()
}

func Menu(s *Stack) {
	var menu int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	data := scanner.Text()
	menu, fallo := strconv.Atoi(data)
	if fallo != nil {
		fmt.Println("ERROR 1: El dato ingresado debe ser un numero entero")
	}
	switch menu {
	case 1:
		fmt.Println("--------1: Ingresar una ecuacion--------\n")
		fmt.Print("Digite la ecuacion en posfijo, separando por espacios: ")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			data := scanner.Text()
			result := InsertarPila(data)
			fmt.Println(fmt.Sprint("Para la ecuacion en posfijo: ", data))
			fmt.Print("La ecuacion en infijo es: ")
			InOrden(result)
			fmt.Println("")
			fmt.Println(fmt.Sprint("El resultado es: ", Operacion(result)))
			break
		}
	case 2:
		fmt.Println("--------2: Salir del programa--------\n")
		os.Exit(2)
	default:
		fmt.Println("ERROR 2: La opcion ingresada en el MENU no es valida")
	}
}

func main() {
	stack := NewStack()
	for {
		fmt.Println("\n...RESOLUCION DE ECUACIONES EN POSFIJO...\n")
		fmt.Println("------------------MENU------------------\n")
		fmt.Println("Seleccione una opcion:\n")
		fmt.Println("1: Ingresar una ecuacion")
		fmt.Println("2: Salir del programa\n")
		Menu(stack)
	}
}
