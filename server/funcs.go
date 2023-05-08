package server

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet{
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Metodo no encontrado \n")
		return
	}
	fmt.Fprintf(w, "Hola %s \n", "PEPE")
}

func getPaises(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(paises)
}

func addPaises(w http.ResponseWriter, r *http.Request) {
	pais := &Pais{}
	err := json.NewDecoder(r.Body).Decode(pais)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v \n", err)
		return
	}

	paises = append(paises, pais)
	fmt.Fprintf(w, "el Pais fue agregado correctamente \n")
}