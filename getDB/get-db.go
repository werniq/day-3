package getDB

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"gorm.io/driver/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Matwyenko1_"
	dbname   = "postgres"
)

// for storinguser details
type User struct {
	gorm.Model
	Name     string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"'`
	Role     string `json:"role"`
}

// Login data
type Authentication struct {
	Email    string `json:"email"`
	password string `json:"password"`
}

// Token is for storing information for correct login credentials
type Token struct {
	Role        string `json:"role"`
	Email       string `json:"email"`
	TokenString string `json:'token'`
}

func getDatabase() *gorm.DB {
	dsn := fmt.Sprint("host=%s port=%s name=%s user=%s password=%s", host, port, dbname, user, password)
	db, err := gorm.Open(postgres.Open(dsn), *gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db
}
