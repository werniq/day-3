package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var (
	db  *gorm.DB
	err error
)

// Model
type Album struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
}

func init() {
	var (
		host     = getEnvVariable("DB_HOST")
		port     = getEnvVariable("DB_PORT")
		user     = getEnvVariable("DB_USER")
		dbname   = getEnvVariable("DB_NAME")
		password = getEnvVariable("DB_Password")
	)

	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		dbname,
		password,
	)
	db, err := gorm.Open("postgres", conn)
	db.AutoMigrate(Album{})

	if err != nil {
		log.Fatal(err)
	}
}

func getEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env variable ", err)
	}
	return os.Getenv(key)
}

func getAlbum(w http.ResponseWriter, r *http.Request) {
	var album []*Album
	// mux.Vars -> to read album id
	id := mux.Vars(r)["id"]
	db.First(&album, id)
	if album.ID == 0 {
		json.NewEncoder(w).Encode("Album not found!")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(album)
}

// Response writer -- Created response after request. Request coming into handler(contains parameters andbody of request)
func postAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") //Setting response headers to .json
	var newAlbum Album                                 // data will be assigned to this var
	json.NewDecoder(r.Body).Decode(&newAlbum)
	db.Create(&newAlbum)                // Write newAlbum into database
	json.NewEncoder(w).Encode(newAlbum) // return created album
}

func getAlbums(w http.ResponseWriter, r *http.Request) {
	var albums []*Album
	db.Find(&albums)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(albums)
}

func updateAlbum(w http.ResponseWriter, r *http.Request) {
	var album Album
	id := mux.Vars(r)["id"]
	db.First(&album, id)
	if album.ID == 0 {
		json.NewEncoder(w).Encode("Album not found!")
		return
	}
	json.NewDecoder(r.Body).Decode(&album)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(album)
}

func deleteAlbum(w http.ResponseWriter, r *http.Request) {
	var album Album
	id := mux.Vars(r)["id"]
	db.First(&album, id)
	if album.ID == 0 {
		json.NewEncoder(w).Encode("Album not found")
		return
	}
	db.Delete(&album, id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Album deleted successfully")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/home", Home).Methods("GET")
	r.HandleFunc("api/v1/albums", postAlbum).Methods("POST")
	r.HandleFunc("api/v1/albums", getAlbum).Methods("GET")
	r.HandleFunc("api/v1/albums/{id}", updateAlbum).Methods("PUT")
	r.HandleFunc("api/v1/albums/{id}", deleteAlbum).Methods("DELETE")
	r.HandleFunc("api/v1/albums/{id}", getAlbums).Methods("GET")
}
