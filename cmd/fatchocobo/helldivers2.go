package main

import (
	"net/http"
)

const HELLDIVERS_URL string = "https://helldiverstrainingmanual.com/api/v1"

type PlanetsResponse map[string]*PlanetData

type PlanetData struct {
	Name string `json:"name"`
	Sector string `json:"sector"`
	Biome string `json:"biome"`
	Environmentals []EnvironmentalData `json:"environmentals"`
}

type EnvironmentalData struct {
	Name string `json:"name"`
	Description string `json:"description"`
}

func GetPlanets() (*http.Response, error) {
	return nil, nil
}
