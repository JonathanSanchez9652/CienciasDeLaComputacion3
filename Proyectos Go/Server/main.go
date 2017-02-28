// Server project main.go
package main

import (
	"net/http"
)

func main() {
	/*
		NO ES BUENO PARA LA PRIVACIDAD PORQUE SE PUEDEN ACCEDER A TODOS
		LOS ARCHIVOS Y HASTA MIRAR SU CÓDIGO
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "index.html")
			//para acceder a todos los archivos de la carpeta
			//http.ServeFile(w, r, r.URL.Path[1:])
		})*/

	//solo es visible lo de la carpeta public
	fileServer := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fileServer))
	//inicia el servidor
	http.ListenAndServe(":8000", nil)
}
