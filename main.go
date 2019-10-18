package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-api-by-generator/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xo/dburl"
	"log"
	"net/http"
	"os"
	"time"
)

var db *sql.DB

func main() {
	Init()
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/add", AddHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

func Init() {
	var err error
	if db, err = DBConnect(); err != nil {
		panic(err.Error())
	}
}

func DBConnect() (*sql.DB, error) {
	dbDriver := os.Getenv("DB_DRIVER")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUrl := fmt.Sprintf("%s://%s:%s@%s/%s?parseTime=true", dbDriver, dbUser, dbPass, dbHost, dbName)
	return dburl.Open(dbUrl)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var user *models.User
	var res []byte
	var id = uint64(22)
	if user, err = models.UserByID(db, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err = json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(res)
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	user := models.User{
		Name: sql.NullString{String: "koba", Valid: true},
		Age: 32,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := user.Insert(db); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = fmt.Fprintf(w, "Add %v", user)
}
