package main

import (
	"log"
	"golang-web-api/internal/routes"
)

func main() {

	router := routes.SetupRouter()

	log.Println("Servidor iniciado na porta 8080")
	
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
