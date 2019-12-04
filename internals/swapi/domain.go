package swapi

type Planet struct {
	Id string `json:"id" bson:"_id"`
	Name string `json:"name"`
	Climate string `json:"climate"`
	Terrain string `json:"terrain"`
	Population string `json:"population"`
}

