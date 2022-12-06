package blog

import (
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/gorilla/mux"
)

// Gorilla/mux + gorm + godotenv

func getDotEnvVar(key string) string {
	err := godotenv.Load(key)
	if err != nil {
		log.Fatalln(err)
	}
	return os.Getenv(key)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("", getDotEnvVar("a"))
}

// Id -> Category
// Id -> Post
// Id -> Post
// Category -> Posts
// UserId -> User (Profile)
// Profile -> Comments
// Profile -> Likes -> Comments
