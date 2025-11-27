package main

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

/*
	tag json để thông báo rằng khi chuyển từ struct
	sang json hay ngược lại sẽ dùng trường json
	vd ID thành id
*/
type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}



func main(){
	r := mux.NewRouter() //khoi tao dich vu gorilla-mux

	//đăng ký đường dẫn: 
	//r.HandleFunc("/đường-dẫn", tên-hàm).Methods("GET" hoặc "POST" hoặc "PUT" hoặc "DELETE")
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe("8000", r)) 

	movies = append(movies,
    Movie{ID: "1", Isbn: "438227", Title: "Inception", Director: &Director{Firstname: "Christopher", Lastname: "Nolan"}},
    Movie{ID: "2", Isbn: "448227", Title: "Parasite", Director: &Director{Firstname: "Bong", Lastname: "Joon-ho"}},
    Movie{ID: "3", Isbn: "123456", Title: "The Matrix", Director: &Director{Firstname: "Lana", Lastname: "Wachowski"}},
	)


}


