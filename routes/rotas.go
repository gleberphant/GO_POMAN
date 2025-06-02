package routes

import (
	"fmt"
	"net/http"
	"poman/controller"
)

type Rotas struct {
}

func CarregarRotas() {
	fmt.Println("Carregar Rotas")

	http.HandleFunc("/", controller.Index)
}
