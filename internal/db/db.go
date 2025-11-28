package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/oita-u/campus-api/internal/config"
)

var Conn *sqlx.DB

func Init() {
	var err error

	Conn, err = sqlx.Connect("postgres", config.C.DBUrl)
	if err != nil {
		log.Fatalf("db connect error: %v", err)
	}
}
