package handlers

import (
	"log"
	"golang-web-api/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GitHubRequest struct {
	Username string `json:"username" binding:"required"`
}

func GetGitHubRepositories(c *gin.Context) {

	var req GitHubRequest
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'username' é obrigatório"})
		return
	}

	data, err := services.FetchGitHubData(req.Username);

	if err != nil {
		log.Printf("Erro ao buscar dados do GitHub: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao consultar o GitHub"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"repositories": data})
}
