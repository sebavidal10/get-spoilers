package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type spoiler struct {
    ID     		string  `json:"id"`
    Content	  string  `json:"content"`
    Movie 		string  `json:"movie"`
}

var spoilers = []spoiler{
	{
		ID: "1",
		Content: "El protagonista es un fantasma",
		Movie: "Sexto Sentido",
	},
	{
		ID: "2",
		Content: "El comisionado gordon no esta muerto",
		Movie: "Dark Knight Rises",
	},
	{
		ID: "3",
		Content: "Cypher es malo",
		Movie: "Matrix",
	},
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func getAllSpoilers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(spoilers)
}

func getSpoilerById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range spoilers {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&spoiler{})
}

func getByMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if len(params["str"]) < 2 {
		http.Error(w, "Invalid request (need more than 1 char)", http.StatusBadRequest)
		return
	}

	var foundSpoilers []spoiler
	for _, item := range spoilers {
		if strings.Contains( strings.ToLower(item.Movie), strings.ToLower(params["str"])) {
			foundSpoilers = append(foundSpoilers, item)
		}
	}
	json.NewEncoder(w).Encode(foundSpoilers)
}

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/spoilers", getAllSpoilers).Methods("GET")
	r.HandleFunc("/spoilers/{id}", getSpoilerById).Methods("GET")
	r.HandleFunc("/spoilers/movie/{str}", getByMovie).Methods("GET")

	http.ListenAndServe(":8080", r)
}
