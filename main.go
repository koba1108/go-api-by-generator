package main

import (
	"database/sql"
	"fmt"
	"github.com/go-api-by-generator/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xo/dburl"
	"html"
	"log"
	"net/http"
	"os"
	"time"
)

var err error
var db *sql.DB

func main() {
	Init()
	http.HandleFunc("/", IndexHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

func Init() {
	if db, err = DBConnect(); err != nil {
		panic(err.Error())
	}
}

func DBConnect() (*sql.DB, error) {
	dbDriver := os.Getenv("DB_DRIVER")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")
	dbUrl := fmt.Sprintf("%s://%s:%s@db/%s", dbDriver, dbUser, dbPass, dbName)
	return dburl.Open(dbUrl)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	user := models.User{
		Name: sql.NullString{String: "koba", Valid: true},
		Age: 32,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err = user.Insert(db); err != nil {
		panic(err.Error())
	}
	_, _ = fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
