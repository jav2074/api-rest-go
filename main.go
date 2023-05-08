package main

import (
	"fmt"

	"api-rest-go/server"
)

func main() {
	fmt.Println("HOLA MANOLA!")
	// log.Fatal("ERROR")

	srv := server.New(":8082")

	err := srv.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

// curl --request GET 127.0.0.1:8082/paises
// curl --request POST 127.0.0.1:8082/paises --data-raw '{"Nombre":"ARG","Lenguaje":"Castellano"}'
//  go run main.go
//  go build main.go
