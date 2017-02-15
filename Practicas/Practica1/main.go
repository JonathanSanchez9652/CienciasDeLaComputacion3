package main

import (
	"fmt"
)

type Node struct {
	Value           int
	Nombre          string
	SegundoNombre   string
	PrimerApellido  string
	SegundoApellido string
	Identificacion  string
	Sintomas        string
	HoraDeLlegada   string
	Eps             string
}

func (n *Node) String() string {
	return fmt.Sprint("->", n.Nombre, "->", n.SegundoNombre, "->", n.PrimerApellido, "->", n.SegundoApellido, "->", n.Identificacion, "->", n.Sintomas, "->", n.HoraDeLlegada, "->", n.Eps)
}

// NewQueue returns a new queue with the given initial size.
func NewQueue(size int) *Queue {
	return &Queue{
		nodes: make([]*Node, size),
		size:  size,
	}
}

// Queue is a basic FIFO queue based on a circular list that resizes as needed.
type Queue struct {
	nodes []*Node
	size  int
	head  int
	tail  int
	count int
}

// Push adds a node to the queue.
func (q *Queue) Push(n *Node) {
	if q.head == q.tail && q.count > 0 {
		nodes := make([]*Node, len(q.nodes)+q.size)
		copy(nodes, q.nodes[q.head:])
		copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])
		q.head = 0
		q.tail = len(q.nodes)
		q.nodes = nodes
	}
	q.nodes[q.tail] = n
	q.tail = (q.tail + 1) % len(q.nodes)
	q.count++
}

// Pop removes and returns a node from the queue in first to last order.
func (q *Queue) Pop() *Node {
	if q.count == 0 {
		return nil
	}
	node := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes)
	q.count--
	return node
}

func main() {

	//fmt.Println("Cola")
	q := NewQueue(1)

	var primerNombre, segundoNombre, identificacion, primerApellido, segundoApellido, sintomas, horaLlegada, eps string

	//nombre slice := make(tipo,tamaño,capacidad-->OPCIONAL)
	slice := make([]string, 0, 20)

	h := 0
	var terminar string
	for {
		fmt.Println("*******************************************************")
		fmt.Println("POR FAVOR A CONTINUACIÓN INGRESE DATOS DE PACIENTE:")

		fmt.Print("Primer nombre: ")
		fmt.Scan(&primerNombre)
		fmt.Print("Segundo nombre: ")
		fmt.Scan(&segundoNombre)
		fmt.Print("Primer apellido: ")
		fmt.Scan(&primerApellido)
		fmt.Print("Segundo apellido: ")
		fmt.Scan(&segundoApellido)
		fmt.Print("Identificacion: ")
		fmt.Scan(&identificacion)
		fmt.Print("Sintomas: ")
		fmt.Scan(&sintomas)
		fmt.Print("Hora de llegada 'hh:mm' : ")
		fmt.Scan(&horaLlegada)
		fmt.Print("EPS: ")
		fmt.Scan(&eps)
		q.Push(&Node{len(q.nodes), primerNombre, segundoNombre, primerApellido, segundoApellido, identificacion, sintomas, horaLlegada, eps})
		h++

		//agregando un nuevo elemento
		slice = append(slice, eps) //--> agrega el elemento 2

		if h > 0 {
			fmt.Print("Finalizar S/N?: ")
			fmt.Scan(&terminar)
			if terminar == "S" {
				fmt.Println("*******************************************************")
				fmt.Println("Numero de pacientes atendidos: ", len(q.nodes)) //Tamaño cola
				break
			} else {
				h = 0
			}

		}
	}
	fmt.Println("***********************************************************")
	fmt.Println("PACIENTES ATENTIDOS POR CADA EPS")
	fmt.Println("Nombre EPS -> No. Pacientes")
	fmt.Println("-----------------------------------------------")
	//IMPRIMIR CUÁNTOS PACIENTES ATENDIÓ CADA EPS
	var palabra string
	var contador int
	var estaRepetida bool
	sliceSinRepetidas := make([]string, 1, 20)

	for i := 0; i < len(slice); i++ {
		palabra = slice[i]

		for k := 0; k < len(sliceSinRepetidas); k++ {
			if sliceSinRepetidas[k] == palabra {
				estaRepetida = true
				break
			}
			estaRepetida = false
		}

		if estaRepetida == false {
			for j := 0; j < len(slice); j++ {
				if slice[j] == palabra {
					contador++
				}
			}
			sliceSinRepetidas = append(sliceSinRepetidas, palabra)
		} else {
			i++
			continue
		}

		fmt.Println(palabra, "->", contador)
		contador = 0

		estaRepetida = false

	}

	//fmt.Println("Slice: ", slice)
	//fmt.Println("Slice sin repetidas: ", sliceSinRepetidas)
	//fmt.Println("Lista EPS: ", slice)
	fmt.Println("***************************************************************************************")
	//Imprimir los pacientes
	fmt.Println("LISTA DE PACIENTES ATENDIDOS")
	fmt.Println("ID->NOMBRE1->NOMBRE2->APELLIDO1->APELLIDO2->IDENTIFICACION->SINTOMAS->HORALLEGADA->EPS")
	fmt.Println("---------------------------------------------------------------------------------------")
	for w := 0; w < len(q.nodes); w++ {
		fmt.Println(w, q.Pop())
	}

}
