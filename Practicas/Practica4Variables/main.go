// Practica4Variables project main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Arbol struct {
	Izquierdo *Arbol
	Valor     string
	Derecho   *Arbol
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

type Variables struct {
	Ecuacion string
	Valor    string
	Variable string
}

type StackVariables struct {
	stackV []*Variables
	count  int
}

func NewStackVariables() *StackVariables {
	return &StackVariables{}
}

func (s *StackVariables) PushV(n *Variables) {
	s.stackV = append(s.stackV[:s.count], n)
	s.count++
}

func (s *StackVariables) PopV() *Variables {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.stackV[s.count]
}

func Operacion(a *Arbol) int {
	if a == nil {
		//fmt.Println("No hay operacion a realizar")
		return 0
	} else if a.Valor == "+" {
		//fmt.Println("Suma")
		return Operacion(a.Izquierdo) + Operacion(a.Derecho)
	} else if a.Valor == "-" {
		//fmt.Println("Resta")
		return Operacion(a.Izquierdo) - Operacion(a.Derecho)
	} else if a.Valor == "*" {
		//fmt.Println("Multiplicacion")
		return Operacion(a.Izquierdo) * Operacion(a.Derecho)
	} else if a.Valor == "/" {
		//fmt.Println("Division")
		return Operacion(a.Izquierdo) / Operacion(a.Derecho)
	} else { //if a.Izquierdo == nil && a.Derecho == nil{
		conv, _ := strconv.Atoi(a.Valor)
		return conv
	}
}

func InsertarPila(a string) *Arbol {
	stack := NewStack()
	val := strings.Split(a, " ")
	for i := 0; i < len(val); i++ {
		arbol := &Arbol{nil, val[i], nil}
		if arbol.Valor != "+" && arbol.Valor != "-" && arbol.Valor != "*" && arbol.Valor != "/" {
			stack.Push(arbol)
		} else {
			arbol.Derecho = stack.Pop()
			arbol.Izquierdo = stack.Pop()
			stack.Push(arbol)
		}
	}
	return stack.Pop()
}

func EncontrarVariable(cadenaCompleta string) ([]string, []string, string) {
	s1 := cadenaCompleta
	if last := len(s1) - 1; last >= 0 && s1[last] == '=' {
		s1 = s1[:last]
	}
	if last := len(s1) - 1; last >= 0 && s1[last] == ' ' {
		s1 = s1[:last]
	}
	if last := len(s1) - 1; last >= 0 && s1[last] == ':' {
		s1 = s1[:last]
	}
	if last := len(s1) - 1; last >= 0 && s1[last] == ' ' {
		s1 = s1[:last]
	}
	arr := strings.Split(s1, " ")
	variables := []string{}
	cadenaTotal := []string{}
	variableNueva := arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	for i := 0; i < len(arr); i++ {
		if arr[i] != "+" && arr[i] != "-" && arr[i] != "*" && arr[i] != "/" {
			if _, err := strconv.Atoi(arr[i]); err == nil {
			} else {
				variables = append(variables, arr[i])
			}
		}
	}
	for i := 0; i < len(arr); i++ {
		cadenaTotal = append(cadenaTotal, arr[i])
	}
	return cadenaTotal, variables, variableNueva
}

func (s *StackVariables) VerVariables() {
	for i := 0; i < s.count; i++ {
		fmt.Println(fmt.Sprint("Ecuacion:", s.stackV[i].Ecuacion))
		fmt.Println(fmt.Sprint("Variable: ", s.stackV[i].Variable))
		fmt.Println(fmt.Sprint("Valor: ", s.stackV[i].Valor))
		fmt.Println("\n")
	}
}

func (s *StackVariables) ValorVar(names []string) []string {
	num := []string{}
	for i := 0; i < len(names); i++ {
		for j := 0; j < s.count; j++ {
			if s.stackV[j].Variable == names[i] {
				num = append(num, s.stackV[j].Valor)
			}
		}
	}
	return num
}

func IntercambiarEcuacion(ecuacion []string, valores []string, variables []string) string {
	ecuacionFinal := ""
	for i := 0; i < len(ecuacion); i++ {
		ecuacionFinal += ecuacion[i]
		ecuacionFinal += " "
	}
	if final := len(ecuacionFinal) - 1; final >= 0 && ecuacionFinal[final] == ' ' {
		ecuacionFinal = ecuacionFinal[:final]
	}
	result := ecuacionFinal
	for i := 0; i < len(valores); i++ {
		result = strings.Replace(result, variables[i], valores[i], -1)
	}
	return result
}

func Menu(s *StackVariables) {
	var menu int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	data := scanner.Text()
	menu, fallo := strconv.Atoi(data)
	if fallo != nil {
		fmt.Println(fmt.Sprint("ERROR 1: El dato ingresado debe ser un numero entero", data))
	}
	switch menu {
	case 1:
		fmt.Println("--------1: Ingresar una ecuacion--------\n")
		fmt.Println("Digite la ecuacion en posfijo, separando por espacios:\n")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			data := scanner.Text()
			ecuacionNew, variables, variableNew := EncontrarVariable(data)
			ecuacion := s.ValorVar(variables)
			ecuacionFinal := IntercambiarEcuacion(ecuacionNew, ecuacion, variables)

			result := InsertarPila(ecuacionFinal)
			fmt.Println(fmt.Sprint("\nPara la ecuacion en posfijo: ", data))
			fmt.Println(fmt.Sprint("\nEL resultado es: ", Operacion(result)))
			valorFinal := strconv.Itoa(Operacion(result))
			s.PushV(&Variables{ecuacionFinal, valorFinal, variableNew})
			break
		}
	case 2:
		fmt.Println("--------2: Ver lista de Variables--------\n")
		s.VerVariables()
	case 3:
		fmt.Println("--------3: Salir del programa--------\n")
		os.Exit(3)
	default:
		fmt.Println("ERROR 2: La opcion ingresada en el MENU no es valida")
	}
}

func main() {
	stack := NewStackVariables()
	for {
		fmt.Println("\n...RESOLUCION DE ECUACIONES EN POSFIJO...\n")
		fmt.Println("------------------MENU------------------\n")
		fmt.Println("Seleccione una opcion:\n")
		fmt.Println("1: Ingresar una ecuacion")
		fmt.Println("2: Ver lista de Variables")
		fmt.Println("3: Salir del programa\n")
		Menu(stack)
	}
}
