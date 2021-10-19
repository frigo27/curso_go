package main

import (
	"net/http"

	"github.com/frigo27/curso_go/routes"
)

func main() {

	routes.CarregaRotas()

	http.ListenAndServe(":8000", nil)
}
