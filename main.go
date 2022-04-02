package main

import (
	"github.com/lucchesisp/go-gin-api/database"
	"github.com/lucchesisp/go-gin-api/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
