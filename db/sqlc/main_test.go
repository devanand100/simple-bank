package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	databaseUrl := "postgresql://postgres:123456@localhost:8000/simple-bank?sslmode=disable"
	conn, err := pgx.Connect(context.Background(), databaseUrl)

	if err != nil {
		log.Fatal("Database connection error")
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
