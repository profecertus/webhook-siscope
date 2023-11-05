package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/webhook-gh", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			fmt.Fprintf(w, "Hola desde el servidor Go!")
		case http.MethodPost:
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusInternalServerError)
				return
			}

			fmt.Fprintf(w, "solicitud post recibida\n")
			fmt.Fprintf(w, "Cuerpo de la solicitud %s\n", string(body))
			fmt.Println(string(body))
		default:
			http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		}
	})

	fmt.Printf("Servidor escuchando en el puerto 8090...\n")
	http.ListenAndServe(":8090", nil)
}
