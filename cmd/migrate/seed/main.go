package main

import (
	"log"

	"github.com/RobTov/habibi-salon/internal/db"
	"github.com/RobTov/habibi-salon/internal/env"
	"github.com/RobTov/habibi-salon/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://postgres:postgres@localhost:5498/postgres?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewPostgresStorage(conn)

	db.Seed(store)
}
