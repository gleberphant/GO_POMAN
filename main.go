package main

import (
	"fmt"
	"net/http"
	"poman/routes"
)

func main() {

	routes.CarregarRotas()

	fmt.Println("Iniciando servidor")
	http.ListenAndServe(":8080", nil)

}
