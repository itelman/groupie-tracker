package models

import (
	"encoding/json"
	"functions/internal/api"
	"strconv"
	"os"
)

const MAP_API_KEY = os.Getenv("GOOGLE_API_KEY")

type Artist struct {
	ID               int      `json:"id"`
	Image            string   `json:"image"`
	Name             string   `json:"name"`
	Members          []string `json:"members"`
	CreationDate     int      `json:"creationDate"`
	FirstAlbum       string   `json:"firstAlbum"`
	LocationsLink    string   `json:"locations"`
	ConcertDatesLink string   `json:"concertDates"`
	RelationsLink    string   `json:"relations"`

	Locations Locations
	Dates     Dates
	Relations Relations

	MapAPIKey string
}

type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`

	LatLngs map[string][]float64
}

type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relations struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func GetArtistsMap() (map[string]Artist, error) {
	var artists []Artist
	body, err := api.ParseJson("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &artists); err != nil {
		return nil, err
	}

	artistsMap := make(map[string]Artist)
	for _, artist := range artists {
		id := strconv.Itoa(artist.ID)
		artistsMap[id] = artist
	}

	return artistsMap, nil
}

func (artist Artist) GetArtistLocations() (Locations, error) {
	var artistLocations Locations
	body, err := api.ParseJson(artist.LocationsLink)
	err = json.Unmarshal(body, &artistLocations)

	return artistLocations, err
}

func (artist Artist) GetArtistDates() (Dates, error) {
	var artistDates Dates
	body, err := api.ParseJson(artist.ConcertDatesLink)
	err = json.Unmarshal(body, &artistDates)

	return artistDates, err
}

func (artist Artist) GetArtistRelations() (Relations, error) {
	var artistRelations Relations
	body, err := api.ParseJson(artist.RelationsLink)
	err = json.Unmarshal(body, &artistRelations)

	return artistRelations, err
}

func (artist *Artist) GetFullInfo() error {
	locations, err := artist.GetArtistLocations()
	if err != nil {
		return err
	}
	artist.Locations = locations

	dates, err := artist.GetArtistDates()
	if err != nil {
		return err
	}
	artist.Dates = dates

	relations, err := artist.GetArtistRelations()
	if err != nil {
		return err
	}
	artist.Relations = relations

	artist.MapAPIKey = MAP_API_KEY

	return nil
}
