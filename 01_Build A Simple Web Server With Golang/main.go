package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"math/rand/v2"
	"github.com/gorilla/mux"
)

/*
tag json để thông báo rằng khi chuyển từ struct
sang json hay ngược lại sẽ dùng trường json
vd ID thành id

json.Encoder để in
json.Decoder để in

*/
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies) //
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {  //r là tham số truyền vào
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r) //nhận r như {id}

	for k, v := range movies{
		if v.ID == param["id"]{ //kiểm tra value xem có = nhau ko, giống như Map Java, get(Key) sẽ ra value
			movies = append(movies[:k], movies[k+1:]...)
			break
		} 
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)

	for _, v := range movies{
		if v.ID == param["id"]{
			json.NewEncoder(w).Encode(v)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie //NewDecoder trả về con chỏ
	_= json.NewDecoder(r.Body).Decode(&movie) //đọc body từ r, lúc này r ko phải là 1 đối tượng Movie, sau đó gán dữ liệu đó cho movie
	
	movie.ID = strconv.Itoa(rand.IntN(1000000))

	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	var updateMovie Movie 
	_ = json.NewDecoder(r.Body).Decode(&updateMovie)

	for index, value := range movies{
		if value.ID == param["id"]{
			updateMovie.ID = param["id"]
			movies[index] = updateMovie
			 
			json.NewEncoder(w).Encode(updateMovie)
			return
		}
	}
}

func main() {
	r := mux.NewRouter() //khoi tao dich vu gorilla-mux

	movies = append(movies,
		Movie{ID: "1", Isbn: "438227", Title: "Inception", Director: &Director{Firstname: "Christopher", Lastname: "Nolan"}},
		Movie{ID: "2", Isbn: "448227", Title: "Parasite", Director: &Director{Firstname: "Bong", Lastname: "Joon-ho"}},
		Movie{ID: "3", Isbn: "123456", Title: "The Matrix", Director: &Director{Firstname: "Lana", Lastname: "Wachowski"}},
	)

	//đăng ký đường dẫn:
	//r.HandleFunc("/đường-dẫn", tên-hàm).Methods("GET" hoặc "POST" hoặc "PUT" hoặc "DELETE")
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))


}
