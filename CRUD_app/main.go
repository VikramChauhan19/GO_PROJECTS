package main
import (
	"fmt"
	"log" // log package is used to print logs (messages) with timestamps
	"github.com/gorilla/mux" //gorilla mux package for routing
	"encoding/json" //encoding and decoding JSON data
	"math/rand" // generating random IDs
	"net/http" // creating web server
	"strconv" // converting int to string
)

type Movie struct {
	ID string `json:"id"` // Movie ID // capital letter makes it publics
	Isbn string `json:"isbn"` // Movie ISBN number
	Title string `json:"title"` // `json:"id"` (Struct Tag)->When converting this struct to JSON, use id instead of ID.”
	Director *Director `json:"director"` // Pointer to Director struct
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`	
}

var movies []Movie // slice to store movies

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json") // Set the content type of the response to application/json
	json.NewEncoder(w).Encode(movies) //Creates a JSON encoder and writes the JSON encoding of movies to the response writer
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)// Get the route parameters
	for i,item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:i],movies[i+1:]...) // ...->spread operator
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie) //Decode needs a address to store the decoded data
	movie.ID = strconv.Itoa(rand.Intn(100000000)) // Generate a random ID for the movie
	movies = append(movies,movie)
	json.NewEncoder(w).Encode(movie)
}
func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json") //👉 This line sets an HTTP response header.
	params := mux.Vars(r)
	for _,item := range movies{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item) //Writing to w is the response , Encoder writes JSON directly into the HTTP response
			return
		}
	}
}

func updateMovie(w http.ResponseWriter, r *http.Request){
	//set json content type
	//params
	//loop through movies , range
	//delete the movie with the id that you have sent
	//add a new movie - the movie that we send in the body of postman
	//encode the movie to response
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i,item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:i],movies[i+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies,movie)
			json.NewEncoder(w).Encode(movie)
		}
	}

}
func main(){
	r := mux.NewRouter()
	movies = append(movies,Movie{ID:"1",Isbn:"438227",Title:"Movie One",Director:&Director{Firstname:"John",Lastname:"Doe"}})
	movies = append(movies,Movie{ID:"2",Isbn:"452555",Title:"Movie Two",Director:&Director{Firstname :"Steve",Lastname:"Smith"}})

	r.HandleFunc("/movies", getMovies).Methods("GET") // GET all movies
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")// GET a single movie by ID
	r.HandleFunc("/movies",createMovie).Methods("POST")// Create a new movie
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")// Update a movie by ID
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")// Delete a movie by ID
	

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080",r))

}