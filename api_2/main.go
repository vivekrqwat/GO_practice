package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id       string    `json:"id"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var movies []Movie

// get movies
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)

}

// delete movies
func deletemovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleting movie")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println("params:", params, params["id"])
	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break

		}
	}
	json.NewEncoder(w).Encode(movies)
}

// get by id
func getMovive(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get by id")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// create body
func createmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication-json")
	var nmv Movie
	_ = json.NewDecoder(r.Body).Decode(&nmv)
	nmv.Id = strconv.Itoa(rand.IntN(1000000))
	movies = append(movies, nmv)
	json.NewEncoder(w).Encode(movies)

}

// update
func updatemovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("update is called")
	params := mux.Vars(r)
	for i, item := range movies {
		if item.Id == params["id"] {
			var up Movie
			json.NewDecoder(r.Body).Decode(&up)
			// item.Id = up.Id
			movies[i].Title = up.Title
			json.NewEncoder(w).Encode(movies[i])
			return
		}

	}
}

func main() {

	M := []Movie{
		{"1", "Inception", &Director{"Christopher", "Nolan"}},
		{"2", "The Matrix", &Director{"Lana", "Wachowski"}},
		{"3", "Interstellar", &Director{"Christopher", "Nolan"}},
	}

	movies = append(movies, M...)
	router := mux.NewRouter()
	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovive).Methods("GET")
	router.HandleFunc("/movies", createmovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updatemovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deletemovie).Methods("DELETE")
	fmt.Println("creating router")
	log.Fatal(http.ListenAndServe(":8080", router))
}
