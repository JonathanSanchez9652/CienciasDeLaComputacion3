// json.go project main.go
package main

import (
	"encoding/json"
	"net/http"
)

type Curso struct {
	Title        string
	NumeroVideos int
}

type Cursos []Curso

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cursos := Cursos{
			Curso{"Curso de go", 30},
			Curso{"Curso de Ruby", 20},
			Curso{"Curso de Java", 50},
		}

		json.NewEncoder(w).Encode(cursos)
	})
	http.ListenAndServe(":8000", nil)
}
