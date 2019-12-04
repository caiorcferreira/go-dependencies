package api

import (
	"context"
	"fmt"
	. "github.com/caiorcferreira/swapi/internals/swapi"
	"github.com/caiorcferreira/swapi/internals/swapi/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type PlanetHandlers struct {
	planetService services.PlanetService
}

func (h PlanetHandlers) GetAllPlanets(c *gin.Context) {
	ctx, _ := context.WithTimeout(c, 1000*time.Millisecond)
	planets, err := h.planetService.GetAll(ctx)
	if err != nil {
		fmt.Printf("An unexpected error occured: %v", err)
		c.JSON(500, gin.H{})
		return
	}

	response := gin.H{
		"count":   len(planets),
		"planets": planets,
	}

	c.JSON(200, response)
}

func (h PlanetHandlers) PostPlanet(c *gin.Context) {
	ctx, _ := context.WithTimeout(c, 1000*time.Millisecond)

	var planet Planet
	if err := c.ShouldBind(&planet); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{})
		return
	}

	saved, err := h.planetService.Create(ctx, planet.Name, planet.Climate, planet.Terrain, planet.Population)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusCreated, saved)
}


func NewPlanetHandlers() PlanetHandlers {
	service := services.NewPlanetService()

	return PlanetHandlers{service}
}

