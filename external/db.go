package external

import (
	"database/sql"
	"github.com/go-api-by-generator/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/martian/log"
	"github.com/xo/dburl"
)

var db *sql.DB
var err error

func initDB() {
	if db, err = dburl.Open(config.DBConnectURL()); err != nil {
		panic(err.Error())
	}
}

func closeDB() {
	err = db.Close()
	if err != nil {
		log.Errorf("db.Close Error: %s", err.Error())
	}
}

func GetDB() *sql.DB {
	return db
}
