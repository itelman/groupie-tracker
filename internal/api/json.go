package api

import (
	"encoding/json"
	"functions/internal/models"
	"io/ioutil"
	"net/http"
	"strconv"
)

func ParseJson(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// fmt.Printf("%#v\n", string(body))

	return body, nil
}

func GetArtistsMap() (map[string]models.Artist, error) {
	var artists []models.Artist
	body, err := ParseJson("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &artists); err != nil {
		return nil, err
	}

	artistsMap := make(map[string]models.Artist)
	for _, artist := range artists {
		id := strconv.Itoa(artist.ID)
		artistsMap[id] = artist
	}

	return artistsMap, nil
}

func GetArtistLocations(artist models.Artist) (models.Locations, error) {
	var artistLocations models.Locations
	body, err := ParseJson(artist.Locations)
	err = json.Unmarshal(body, &artistLocations)

	return artistLocations, err
}

func GetArtistDates(artist models.Artist) (models.Dates, error) {
	var artistDates models.Dates
	body, err := ParseJson(artist.ConcertDates)
	err = json.Unmarshal(body, &artistDates)

	return artistDates, err
}

func GetArtistRelations(artist models.Artist) (models.Relations, error) {
	var artistRelations models.Relations
	body, err := ParseJson(artist.Relations)
	err = json.Unmarshal(body, &artistRelations)

	return artistRelations, err
}

func GetArtistFullInfo(artist models.Artist) (models.ArtistInfo, error) {
	emptyStruct := models.ArtistInfo{}

	locations, err := GetArtistLocations(artist)
	if err != nil {
		return emptyStruct, err
	}

	dates, err := GetArtistDates(artist)
	if err != nil {
		return emptyStruct, err
	}

	relations, err := GetArtistRelations(artist)
	if err != nil {
		return emptyStruct, err
	}

	return models.ArtistInfo{BasicInfo: artist, Locations: locations, Dates: dates, Relations: relations}, nil
}
