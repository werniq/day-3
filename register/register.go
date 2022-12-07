package main

import (
	"blog"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"gorm.io/driver/postgres"
	"html/template"
	"net/http"
)

// For db
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Matwyenko1_"
	dbname   = "postgres"
)

var tpl *template.Template
var db *gorm.DB

func main() {
	tpl, _ = template.ParseGlob("./template/register.gohtml")
	dsn := fmt.Sprint("host=%s port=%s name=%s user=%s password=%s", host, port, dbname, user, password)

	// Opens database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// Closes database
	defer db.Close()
	r := mux.NewRouter()

	r.HandleFunc("/register", registerHandler).Methods("POST")
	r.HandleFunc("/registerauth", registerAuthHandler)
	r.HandleFunc("/registererror", registerError).Methods("GET")

	fmt.Println("Database connected - #Register")
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	// Post request -> username & password
	var username string
	var userPass string
	var user *blog.User
	if user == nil {
		fmt.Println("hahahah")
	}
	if !(r.Method == "POST" && r.URL.Path == "/register") {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	r.ParseForm()
	username = r.FormValue("username")
	userPass = r.FormValue("password")
	if username != "" {
		if len(username) < 10 {
			user.Username = username
		} else {
			fmt.Println("<h1>Length of username must be more than 10!</h1>")
		}
		if userPassword == "" {

		} else {
			http.Redirect(w, r, "/registererror", 0)
		}
	} else {
		http.Redirect(w, r, "/registererror", 0)
	}
	//res := db.Create(&user)
}

func registerAuthHandler(w http.ResponseWriter, r *http.Request) {

}

func registerError(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseGlob("./template/no-data.gohtml")

	if err == nil {
		panic(err)
	}

	templ.Execute(w, nil)
}
