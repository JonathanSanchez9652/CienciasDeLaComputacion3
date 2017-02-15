// GoChannels project main.go
package main

import (
	"fmt"
)

func main() {
	channel := make(chan string)

	go func(channel chan string) {
		for {
			var name string
			fmt.Scanln(&name)
			channel <- name //cuando envíamos
		}
	}(channel)

	msg := <-channel //cuando recibimos
	fmt.Println("Primer mensaje recibido en el canal: ", msg)

	msg = <-channel //cuando recibimos
	fmt.Println("Segundo Mensaje recibido en el canal: ", msg)
}
