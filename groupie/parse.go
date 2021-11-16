package student

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

// "MainPage": Handles page with info about all Artits
func MainPage(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" && r.Method == "GET" {
		Err("404 Not Found", http.StatusNotFound, w)
		return
	}

	if r.URL.Path == "/" && r.Method != "GET" {
		Err("405 Method Not Allowed", http.StatusMethodNotAllowed, w)
		return
	}

	htmlTemplate, err := template.ParseFiles("templates/groupie.html")

	if err != nil {
		Err("500 Internal Server Error", http.StatusInternalServerError, w)
		return
	}

	GettingAPIData(w)

	htmlTemplate.Execute(w, ArtistsInfo)
}

// "ArtistPage": Handles page with precise data about Artists
func ArtistPage(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Path) < 10 || (r.URL.Path[:9] != "/artists/") {
		Err("404 Not Found", http.StatusNotFound, w)
		return
	}

	if r.Method != "GET" {
		Err("405 Method Not Allowed", http.StatusMethodNotAllowed, w)
		return
	}

	htmlTemplate, err := template.ParseFiles("templates/artist.html")

	if err != nil {
		Err("500 Internal Server Error", http.StatusInternalServerError, w)
		return
	}

	ArtistID := strings.TrimPrefix(r.URL.Path, "/artists/")

	ID, err := strconv.Atoi(ArtistID)

	if err != nil {
		Err("400 Bad Request", http.StatusBadRequest, w)
		return
	}

	GettingAPIData(w)

	if len(ArtistsInfo) < ID {
		Err("404 Not Found", http.StatusNotFound, w)
		return
	} else if ID < 1 {
		Err("400 Bad Request", http.StatusBadRequest, w)
		return
	}

	err = htmlTemplate.Execute(w, ArtistsInfo[ID-1])

	if err != nil {
		Err("500 Internal Server Error", http.StatusInternalServerError, w)
		return
	}
}

// "MainHandler": Function for calling handle functions.
func MainHandler() {
	val := http.FileServer(http.Dir("style"))
	http.Handle("/style/", http.StripPrefix("/style", val))
	http.HandleFunc("/", MainPage)
	http.HandleFunc("/artists/", ArtistPage)
	http.ListenAndServe(":7770", nil)
}
