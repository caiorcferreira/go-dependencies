package api

import (
	"github.com/caiorcferreira/swapi/internals/swapi/services"
	"github.com/gin-gonic/gin"
)

type PlanetHandlers struct {
	planetService services.PlanetService
}

func (p PlanetHandlers) GetAllPlanets(c *gin.Context) {
	planets, err := p.planetService.GetAll()
	if err != nil {
		c.JSON(500, nil)
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

