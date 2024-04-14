package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

const HELLDIVERS2_URL string = "https://helldiverstrainingmanual.com/api/v1"

type WarCampaignResponse []struct {
	PlanetIndex int     `json:"planetIndex"`
	Name        string  `json:"name"`
	Faction     string  `json:"faction"`
	Players     int     `json:"players"`
	Health      int     `json:"health"`
	MaxHealth   int     `json:"maxHealth"`
	Percentage  float32 `json:"percentage"`
	Defense     bool    `json:"defense"`
	MajorOrder  bool    `json:"majorOrder"`
	Biome       struct {
		Slug        string `json:"slug"`
		Description string `json:"description"`
	} `json:"biome"`
	ExpireDateTime float32 `json:"expireDateTime"`
}

type PlanetsResponse map[string]*PlanetData

type PlanetData struct {
	Name           string              `json:"name"`
	Sector         string              `json:"sector"`
	Biome          string              `json:"biome"`
	Environmentals []EnvironmentalData `json:"environmentals"`
}

type EnvironmentalData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// API calls
func GetWarCampaign() (*http.Response, error) {
	endpoint := HELLDIVERS2_URL + "/war/campaign"

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)

	if err != nil {
		return nil, err
	}

	return MakeRequest(request)
}

// Extract responses
func GetWarCampaignResponse() *WarCampaignResponse {
	data := new(WarCampaignResponse)

	response, err := GetWarCampaign()

	if err != nil {
		log.Printf("Failed to create HTTP request\n\tError: %s")
		return nil
	} else if response.StatusCode != 200 {
		log.Printf("Bad status code: %s", response.Status)
		return nil
	}

	raw, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		log.Printf("Unable to read body\n\tError: %s", err)
	}

	if err = json.Unmarshal(raw, data); err != nil {
		log.Printf("WARNING: %s", err)
	}

	return data
}

func GetWarCampaignPlanets() []string {
	campaignInfo := GetWarCampaignResponse()
	var planets []string

	for _, i := range *campaignInfo {
		planets = append(planets, i.Name)
	}

	return planets
}
