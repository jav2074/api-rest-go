package server

import (
	"fmt"
	"net/http"
)

func initRutas(){
	http.HandleFunc("/", index)
	http.HandleFunc("/paises", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getPaises(w,r)
		case http.MethodPost:
			addPaises(w,r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Metodo no encontrado \n")
			return
		}
	})
}