package main

import (
	"database/sql"
	"github.com/go-study/rest-api-go/api"
	"github.com/go-study/rest-api-go/db/gorm/models"
	db "github.com/go-study/rest-api-go/db/sqlc"
	_ "github.com/lib/pq"
	"log"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	g := models.Gorm{}
	g.DbInit()
	store := db.NewStore(conn)
	server := api.NewServer(store, &g)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
