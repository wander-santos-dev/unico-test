package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	HOST    = "localhost"
	PORT    = 5432
	SSLMODE = "disable"
)

var ErrNoMatch = fmt.Errorf("No matching record!")

type Database struct {
	Conn *sql.DB
}

func Initialize(username, password, dbname string) (Database, error) {
	db := Database{}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		HOST, PORT, username, password, dbname, SSLMODE)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return db, err
	}

	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}

	log.Println("Database connection established")
	return db, nil
}
