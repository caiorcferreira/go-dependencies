package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRoutes() http.Handler {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"health": true})
	})

	planetHandlers := NewPlanetHandlers()

	planets := router.Group("/planets")
	{
		planets.GET("/", planetHandlers.GetAllPlanets)
	}

	return router
}
