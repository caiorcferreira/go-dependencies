package api

import (
	"context"
	"github.com/caiorcferreira/swapi/internals/swapi/services"
	"github.com/gin-gonic/gin"
	"time"
)

type PlanetHandlers struct {
	planetService services.PlanetService
}

func (p PlanetHandlers) GetAllPlanets(c *gin.Context) {
	ctx, _ := context.WithTimeout(c, 1000*time.Millisecond)
	planets, err := p.planetService.GetAll(ctx)
	if err != nil {
		c.JSON(500, gin.H{})
		return
	}

	response := gin.H{
		"count":   len(planets),
		"planets": planets,
	}

	c.JSON(200, response)
}


func NewPlanetHandlers() PlanetHandlers {
	service := services.NewPlanetService()

	return PlanetHandlers{service}
}

