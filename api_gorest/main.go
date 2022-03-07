package main

import (
	"fmt"

	"github.com/italodasilvaa/go-rest-api/database"
	"github.com/italodasilvaa/go-rest-api/models"
	"github.com/italodasilvaa/go-rest-api/routes"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Italo", Historia: "Texto de teste "},
		{Id: 2, Nome: "Igor", Historia: "Texto de teste "},
		{Id: 3, Nome: "Luana", Historia: "Texto de teste "},
	}
	database.ConectaBancoDeDados()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
