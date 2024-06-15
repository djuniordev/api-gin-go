package main

import (
	"github.com/djuniordev/api-go-gin/database"
	"github.com/djuniordev/api-go-gin/routes"
)

func main() {

	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
