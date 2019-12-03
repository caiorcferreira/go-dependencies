package services

type PlanetService struct {}

type Planet struct {
	Name string `json:"name"`
	Climate string `json:"climate"`
	Terrain string `json:"terrain"`
	Population string `json:"population"`
}

func (p PlanetService) GetAll() ([]Planet, error) {
	return []Planet{
		{
			Name:       "Tatooine",
			Climate:    "arid",
			Terrain:    "desert",
			Population: "200000",
		},
	}, nil
}

func NewPlanetService() PlanetService {
	return PlanetService{}
}
