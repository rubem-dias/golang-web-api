package routes

import (
	"golang-web-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	
	r := gin.Default()

	r.POST("/github", handlers.GetGitHubRepositories)

	return r
}
