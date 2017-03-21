// Practica5Tokens project main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Arbol struct {
	Izquierda *Arbol
	Valor     string
	Derecha   *Arbol
}

func NewArbol() *Arbol {
	return &Arbol{}
}

type Stack struct {
	arbol []*Arbol
	count int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(n *Arbol) {
	s.arbol = append(s.arbol[:s.count], n)
	s.count++
}

func (s *Stack) Pop() *Arbol {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.arbol[s.count]
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
		return 0
	} else if a.Valor == "+" {
		return Operacion(a.Izquierda) + Operacion(a.Derecha)
	} else if a.Valor == "-" {
		return Operacion(a.Izquierda) - Operacion(a.Derecha)
	} else if a.Valor == "*" {
		return Operacion(a.Izquierda) * Operacion(a.Derecha)
	} else if a.Valor == "/" {
		if Operacion(a.Derecha) == 0 {
			fmt.Println("***************************")
			fmt.Println("ERROR DIVISIÓN POR CERO")
			fmt.Println("***************************")
		} else {
			return Operacion(a.Izquierda) / Operacion(a.Derecha)
		}
		return 0
	} else {
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
			arbol.Derecha = stack.Pop()
			arbol.Izquierda = stack.Pop()
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
		fmt.Println(fmt.Sprint("Identificador: ", s.stackV[i].Variable))
		fmt.Println(fmt.Sprint("Resultado: ", s.stackV[i].Valor))
		fmt.Println("\n")
	}
}

func (s *StackVariables) imprimirVariable() string {
	return s.stackV[s.count-1].Variable
}

func (s *StackVariables) verificaVariableValida() bool {
	v := s.stackV[s.count-1].Variable

	var validUno = regexp.MustCompile(`^[A-Z]+[a-z]+[_]+[0-9]+$`)
	veriUno := validUno.MatchString(v)

	var validDos = regexp.MustCompile(`^[A-Z]+[a-z]+$`)
	veriDos := validDos.MatchString(v)

	var validTres = regexp.MustCompile(`^[A-Z]+$`)
	veriTres := validTres.MatchString(v)

	if veriUno == true || veriDos == true || veriTres == true {
		return true
	}
	return false
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

func InOrden(a *Arbol) {
	if a == nil {
		return
	}
	InOrden(a.Izquierda)
	fmt.Printf(a.Valor)
	InOrden(a.Derecha)
}

func Menu(s *StackVariables) {
	var menu int
	//var variableString string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	data := scanner.Text()
	menu, fallo := strconv.Atoi(data)
	if fallo != nil {
		fmt.Println(fmt.Sprint("ERROR 1: El dato ingresado debe ser un numero entero", data))
	}
	switch menu {
	case 1:
		fmt.Println("\n--------1: Ingresar una ecuacion--------\n")
		fmt.Println("Digite la ecuacion en posfijo, separando por espacios")
		fmt.Println("La ecuación debe terminar con el nombre de la variable (Variable inicia con letra en MAYÚSUCULA)")
		fmt.Println("**OPCIONAL: nombreDeVariable + ':=' ****")
		fmt.Print("\n--> Ingrese ecuación: ")
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {

			data := scanner.Text()
			ecuacionNew, variables, variableNew := EncontrarVariable(data)
			ecuacion := s.ValorVar(variables)
			ecuacionFinal := IntercambiarEcuacion(ecuacionNew, ecuacion, variables)
			result := InsertarPila(ecuacionFinal)
			valorFinal := strconv.Itoa(Operacion(result))
			s.PushV(&Variables{ecuacionFinal, valorFinal, variableNew})

			if s.verificaVariableValida() == false {
				fmt.Println("***** VARIABLE INVÁLIDA*********")
				fmt.Println("***** PROGRAMA TERMINADO*********")
				os.Exit(3)
			} else {
				fmt.Println("")
				fmt.Println("****SOLUCIÓN*****  ")
				fmt.Println("")
				fmt.Println(fmt.Sprint("Ecuacion en posfijo: ", data))
				fmt.Print("Ecuacion en infijo: ", s.imprimirVariable(), " := ")
				InOrden(result)
				fmt.Println("\nResultado: ", s.imprimirVariable(), " = ", Operacion(result))
			}
			break
		}
	case 2:
		fmt.Println("--------2: Ver lista de Variables--------\n")
		s.VerVariables()
	case 3:
		fmt.Println("--------3: Salir del programa--------\n")
		fmt.Println("***** PROGRAMA TERMINADO*********")
		os.Exit(3)
	default:
		fmt.Println("ERROR 2: La opcion ingresada en el MENU no es valida")
	}
}

func main() {
	stack := NewStackVariables()
	for {
		fmt.Println("\n***ECUACIONES EN POSFIJO****\n")
		fmt.Println("---->Presione la tecla indicada para seleccionar una opción:\n")
		fmt.Println("Ingresar una ecuacion -> 1 ")
		fmt.Println("Ver lista de Variables -> 2 ")
		fmt.Println("Salir del programa -> 3 \n")
		fmt.Print("Ingrese su selección: ")
		Menu(stack)
	}
}
