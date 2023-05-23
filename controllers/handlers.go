package controllers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/abhisenberg/go-movies/models"
	"github.com/gorilla/mux"
)

var Movies []models.Movie

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newMovie models.Movie

	json.NewDecoder(r.Body).Decode(&newMovie)
	newMovie.ID = strconv.Itoa(rand.Intn(1000))

	Movies = append(Movies, newMovie)
	json.NewEncoder(w).Encode(newMovie)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	fmt.Println(vars)
	targetID := vars["id"]
	for _, item := range Movies {
		if item.ID == targetID {
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Movies)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	targetID := mux.Vars(r)["id"]
	for index, item := range Movies {
		if item.ID == targetID {

			var newMovie models.Movie
			json.NewDecoder(r.Body).Decode(&newMovie)
			newMovie.ID = item.ID

			Movies = append(Movies[:index], Movies[index+1:]...)
			Movies = append(Movies, newMovie)

			json.NewEncoder(w).Encode(newMovie)
			break
		}
	}
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	targetID := mux.Vars(r)["id"]
	for index, item := range Movies {
		if item.ID == targetID {
			Movies = append(Movies[:index], Movies[index+1:]...)
			json.NewEncoder(w).Encode(Movies)
			break
		}
	}
}
