package server

import (
	"net/http"
)

type Pais struct{
	Nombre 		string
	Lenguaje 	string
}

var paises []*Pais = nil

func New(addr string) *http.Server{
	initRutas()

	return &http.Server{
		Addr: addr,
	}
}