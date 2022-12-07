package register

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"gorm.io/driver/postgres"
	"html/template"
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

	fmt.Println("Database connected - #Register")
}
