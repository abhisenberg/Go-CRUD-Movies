package main

import (
	"net/http"

	"github.com/abhisenberg/go-movies/controllers"
	"github.com/abhisenberg/go-movies/models"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", controllers.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", controllers.GetMovie).Methods("GET")
	r.HandleFunc("/movies", controllers.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", controllers.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", controllers.DeleteMovie).Methods("DELETE")

	controllers.Movies = append(controllers.Movies, models.Movie{"1", "Shawshank Redemption", &models.Director{"Jon", "Sno"}})
	controllers.Movies = append(controllers.Movies, models.Movie{"2", "Batman Begins", &models.Director{"Chris", "Nolan"}})

	http.ListenAndServe(":8000", r)
}
