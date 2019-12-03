package services

import (
	"context"
	. "github.com/caiorcferreira/swapi/internals/swapi"
	"github.com/caiorcferreira/swapi/internals/swapi/infra"
)

type PlanetService struct {
	planetRepo infra.PlanetRepository
}


func (s PlanetService) GetAll(ctx context.Context) ([]Planet, error) {
	planets, err := s.planetRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return planets, nil
}

func NewPlanetService() PlanetService {
	pr := infra.NewPlanetRepository()

	return PlanetService{pr}
}
