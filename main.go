package blog

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"gorm.io/driver/postgres"
	"net/http"
	"strconv"
)

// For db
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Matwyenko1_"
	dbname   = "postgres"
)

dsn := fmt.Sprint("host=%s port=%s name=%s user=%s password=%s", host, port, dbname, user, password)
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
defer db.Close()

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

// Posts
type Post struct {
	id       int         `json:"id"`
	text     string      `json:"text"`
	title    string      `json:"title"`
	author   *User       `json:"user"`
	comments []*Comment  `json:"comments"`
	likes    int         `json:"likes"`
	category []*Category `json:"categories"`
}

// Comment to posts
type Comment struct {
	author *User  `json:"user"`
	likes  int    `json:"likes"`
	text   string `json:"text"`
}

// For authorization
type User struct {
	id       	  int    `json:"int"`
	Username      string `json:"username"`
	email 	 	  string `json:email`
	password 	  string `json:"password"`
	totalLikes    int    `json:"totalLikes"`
	totalComments int    `json:"totalComments"`
}

// Category for posts
type Category struct {
	business       string `json:"category-business"`
	travel         string `json:"category-travel"`
	cryptocurrency string `json:"category-cryptocurrency"`
	cooking        string `json:"category-cooking"`
	books          string `json:"category-books"`
	art            string `json:"category-art"`
}

// Create Post function
func createPost(w http.ResponseWriter, r *http.Request) {
	var post *Post
	json.NewDecoder(r.Body).Decode(&post)
	postId++
	post.id = postId
	posts = append(posts, post)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

// Get Post by id
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

// Update post
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

// Delete post from db
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

// Get all posts(home page)
func getAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func main() {

	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	fmt.Println("Starting server...")

	r.HandleFunc("/", getAllPosts).Methods("GET")
	r.HandleFunc("/post", createPost).Methods("POST")

	r.HandleFunc("/post/{id}", getPost).Methods("GET")
	r.HandleFunc("/post/{id}", deletePost).Methods("DELETE")
	r.HandleFunc("post/{id}", modifyPost).Methods("PUT")

	r.HandleFunc("/signup", SignUp).Methods("POST")
	r.HandleFunc("/signin", SignIn).Methods("POST")

	fmt.Println(http.ListenAndServe(":8080", nil))
}
