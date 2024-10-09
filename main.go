package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"strconv"
	"github.com/gorilla/mux"
	"math/rand"
)

type Movie struct {
    ID       string   `json:"id"`
    Isbn     string   `json:"isbn"`
    Title    string   `json:"title"`
    Director *Director `json:"director"`
}

type Director struct {
    Firstname string `json:"firstname"`
    Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")
	
	params := mux.Vars(r)
	for index, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")

	params := mux.Vars(r)
	for _, item := range movies{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")

	var movie Movie
	
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
    	http.Error(w, err.Error(), http.StatusBadRequest)
    	return
	}
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
	
}

func updateMovie(w http.ResponseWriter, r *http.Request){
	//set json content type
	w.Header().Set("content type", "application/json")
	//params
	params := mux.Vars(r)
	//Loop over the movies, range
	//delete the movie with the i.d that you've sent
	//pass a movie - the movie that we send in the body of postman
	for index, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)

			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	json.NewEncoder(w).Encode(movies)

}

func main(){
	r := mux.NewRouter()

	movies = append(movies, Movie{ 
		ID: "1", 
		Isbn: "438227", 
		Title: "Movie1", 
		Director: &Director{Firstname: "John", Lastname: "Doe"},
	})
	
	movies = append(movies, Movie{ 
		ID: "2", 
		Isbn: "432313", 
		Title: "Movie2", 
		Director: &Director{Firstname: "Musk", Lastname: "Elon"},
	})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("The server is started at: 8000\n ")
	log.Fatal(http.ListenAndServe(":8000", r))
}
