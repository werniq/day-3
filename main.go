package blog

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"net/http"
	"os"
	"strconv"
)

var (
	postId = 0
	posts  []*Post
	keys   = []string{
		"DB_NAME",
		"DB_HOST",
		"DB_USER",
		"DB_PORT",
		"DB_PASSWORD",
	}
)

type Post struct {
	id       int         `json:"id"`
	text     string      `json:"text"`
	title    string      `json:"title"`
	author   *User       `json:"user"`
	comments []*Comment  `json:"comments"`
	likes    int         `json:"likes"`
	category []*Category `json:"categories"`
}

type Comment struct {
	author *User  `json:"user"`
	likes  int    `json:"likes"`
	text   string `json:"text"`
}

type User struct {
	id            int    `json:"int"`
	username      string `json:"username"`
	password      string `json:"password"`
	totalLikes    int    `json:"totalLikes"`
	totalComments int    `json:"totalComments"`
}

type Category struct {
	business       string `json:"category1"`
	travel         string `json:"category2"`
	cryptocurrency string `json:"category3"`
	cooking        string `json:"category4"`
	books          string `json:"category5"`
	art            string `json:"category6"`
}

func createPost(w http.ResponseWriter, r *http.Request) {
	var post *Post
	json.NewDecoder(r.Body).Decode(&post)
	postId++
	post.id = postId
	posts = append(posts, post)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	for _, post := range posts {
		if id == strconv.Itoa(post.id) {
			json.NewEncoder(w).Encode(post)
			return
		}
	}
}

func getDotEnvVar(key string) string {
	err := godotenv.Load(key)
	if err != nil {
		panic(err)
	}
	return os.Getenv(key)
}

func modifyPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	for i, post := range posts {
		if strconv.Itoa(post.id) == id {
			posts = append(posts[:i], posts[i+1:]...)
			var updatedPost *Post
			json.NewDecoder(r.Body).Decode(&updatedPost)
			posts = append(posts, updatedPost)
			json.NewEncoder(w).Encode(updatedPost)
			return
		}
	}
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	for i, post := range posts {
		if strconv.Itoa(post.id) == id {
			posts = append(posts[:i], posts[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
}

func getAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func main() {
	dsn := "host=localhost user=postgres password=Matwyenko1_ dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.DB()
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", getAllPosts).Methods("GET")
	r.HandleFunc("/post", createPost).Methods("POST")

	r.HandleFunc("/post/{id}", getPost).Methods("GET")
	r.HandleFunc("/post/{id}", deletePost).Methods("DELETE")
	r.HandleFunc("post/{id}", modifyPost).Methods("PUT")
}
