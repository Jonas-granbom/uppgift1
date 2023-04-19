package data

type Player struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	TeamId       uint   `json:"team_id"`
	JerseyNumber int    `json:"jersey_number"`
	BirthYear    int    `json:"birth_year"`
}

type Team struct {
	Id          int    `json:"team_id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	FoundedYear int    `json:"founded_year"`
}
