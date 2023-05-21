package main

import (
	"fmt"
	"github.com/mthsnts/go-API/database"
	"github.com/mthsnts/go-API/models"
	"github.com/mthsnts/go-API/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
