package db

import (
    "database/sql"
    "log"

    _ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func Connect() {
    var err error
    connStr := "postgres://kendir:kendir@localhost:5432/kendir?sslmode=disable"
    DB, err = sql.Open("pgx", connStr)
    if err != nil {
        log.Fatal("DB open error:", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatal("Cannot connect to DB:", err)
    }

    log.Println("Connected to Database!")
}
