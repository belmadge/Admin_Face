package main

import (
	"log"
	"os"

	db "github.com/belmadge/Admin_Face/infra/repository"
	"github.com/belmadge/Admin_Face/routes"
	"github.com/belmadge/Admin_Face/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Carregar variáveis de ambiente
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Conectar ao banco de dados
	db.ConnectDatabase()

	// Configuração do gin
	router := gin.Default()

	// Middleware para registrar atividades
	router.Use(utils.ActivityLogger())

	// Inicializar rotas
	routes.InitializeRoutes(router)

	// Iniciar servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	router.Run(":" + port)
}
