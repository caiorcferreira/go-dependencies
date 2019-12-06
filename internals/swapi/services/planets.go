package services

import (
	"context"
	"fmt"
	. "github.com/caiorcferreira/swapi/internals/swapi"
	"github.com/caiorcferreira/swapi/internals/swapi/infra"
)

type PlanetService struct {
	planetRepo infra.PlanetRepository
	filmsService FilmsService
}

func (s PlanetService) GetAll(ctx context.Context) ([]Planet, error) {
	planets, err := s.planetRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return planets, nil
}

func (s PlanetService) Create(ctx context.Context, name, climate, terrain, population string) (Planet, error) {
	newPlanet := Planet{
		Name:       name,
		Climate:    climate,
		Population: population,
		Terrain:    terrain,
	}

	appearances, err := s.filmsService.GetFilmAppearances(newPlanet)
	if err != nil {
		fmt.Printf("failed to retrieve films information for planet %v due to %v", newPlanet, err)
	}
	newPlanet.Films = appearances

	savedPlanet, err := s.planetRepo.Save(ctx, newPlanet)
	if err != nil {
		return Planet{}, err
	}

	return savedPlanet, nil
}

func NewPlanetService() PlanetService {
	pr := infra.NewPlanetRepository()
	filmsService := NewFilmsService()

	return PlanetService{pr, filmsService}
}
