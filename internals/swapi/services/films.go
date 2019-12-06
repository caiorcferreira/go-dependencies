package services

import (
	"encoding/json"
	"errors"
	"fmt"
	. "github.com/caiorcferreira/swapi/internals/swapi"
	"net/http"
)

type FilmsService struct {}

type planetDto struct {
	Name string `json:"name"`
	Films []string `json:"films"`
}

type searchResultDto struct {
	Results []planetDto `json:"results"`
}

func (s FilmsService) GetFilmAppearances(planet Planet) (int, error) {
	planetName := planet.Name
	url := fmt.Sprintf("https://swapi.co/api/planets/?search=%s", planetName)
	response, err := http.Get(url)
	if err != nil || response.StatusCode != http.StatusOK {
		return 0, err
	}
	defer response.Body.Close()

	var search searchResultDto
	decodeErr := json.NewDecoder(response.Body).Decode(&search)
	if decodeErr != nil {
		return 0, err
	}

	planets := search.Results

	if p, ok := findByName(planets, planetName); ok {
		return len(p.Films), nil
	}

	return 0, errors.New(fmt.Sprintf("no planet found with name %s\n", planetName))
}

func findByName(planets []planetDto, name string) (planetDto, bool) {
	for _, p := range planets {
		if p.Name == name {
			return p, true
		}
	}

	return planetDto{}, false
}

func NewFilmsService() FilmsService {
	return FilmsService{}
}
