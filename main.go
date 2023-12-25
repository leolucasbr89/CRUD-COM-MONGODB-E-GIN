package main

import (
	"mongodb-com-go/routes"
	"mongodb-com-go/database"
)

func main() {
	database.ConexaoComBancoDeDados()
	routes.LidaComRequisicoes()
}