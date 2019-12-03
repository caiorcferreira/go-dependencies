package services

import (
	. "github.com/caiorcferreira/swapi/internals/swapi"
	"github.com/caiorcferreira/swapi/internals/swapi/infra"
)

type PlanetService struct {
	planetRepo infra.PlanetRepository
}


func (s PlanetService) GetAll() ([]Planet, error) {
	planets, err := s.planetRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return planets, nil
}

func NewPlanetService() PlanetService {
	pr := infra.NewPlanetRepository()

	return PlanetService{pr}
}
