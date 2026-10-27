package main

import (
	"context"
	"log"
	"os"
	"time"

	"task-tracker/internal/app"
	"task-tracker/internal/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	app := app.BuildApp()
	defer app.Close()

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("подключение к PostgreSQL работает")

	store := storage.NewPostgresStorage(pool)
	_ = store
}