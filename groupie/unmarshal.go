package student

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// "ArtistsInfo": var for struct Artists (look for 'structs.go').
var ArtistsInfo []Artists

// "RelationInfo": var for struct Relation (look for 'structs.go').
var RelationInfo = Relation{}

// "FillingDatesForArtists": Filling empty part (DatesLocations) of the struct Artists.
func FillingDatesForArtists() {
	for index := range ArtistsInfo {
		ArtistsInfo[index].DatesLocations = RelationInfo.Index[index].DatesLocations
	}
}

// "UnmarshalAPIData": Process of unmarshalling data by using given api.
func UnmarshalAPIData(url string, val interface{}, w http.ResponseWriter) {
	res, err := http.Get(url)
	if err != nil {
		Err("500 Internal Server Error", http.StatusInternalServerError, w)
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		Err("500 Internal Server Error", http.StatusInternalServerError, w)
		return
	}
	err = json.Unmarshal(body, &val)
	if err != nil {
		Err("500 Internal Server Error", http.StatusInternalServerError, w)
		return
	}
}

// "GettingAPIData": Filling appropriate structs with data from api.
func GettingAPIData(w http.ResponseWriter) {
	var ArtitsURL = "https://groupietrackers.herokuapp.com/api/artists"
	var RelationURL = "https://groupietrackers.herokuapp.com/api/relation"

	UnmarshalAPIData(ArtitsURL, &ArtistsInfo, w)
	UnmarshalAPIData(RelationURL, &RelationInfo, w)

	FillingDatesForArtists()
}
