// HolaWeb project main.go
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//forma 1 de responder una petición
	http.HandleFunc("/holamundo", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hoooooola")
	})
	//forma 2 de responder una petición
	http.HandleFunc("/", handler)     //cuando lleguen a la / responde con la función handler
	http.ListenAndServe(":8000", nil) //levanta el servidor en el puerto 8000
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hay una nueva petición")
	io.WriteString(w, "Hola mundo")
}
