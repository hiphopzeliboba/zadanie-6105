package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v4/pgxpool"
	"zadanie-6105/internal/config"
	tender2 "zadanie-6105/internal/service/tender"

	"log"
	"net/http"
	"os"
	"zadanie-6105/internal/api/handlers"
	"zadanie-6105/internal/repository/tender"
)

func main() {
	ctx := context.Background()
	config.Load(".env")

	pool, err := pgxpool.Connect(ctx, os.Getenv("POSTGRES_CONN"))
	if err != nil {
		log.Fatal(err)
	}

	tRepo := tender.NewRepository(pool)

	s := tender2.NewService(tRepo)

	impl := handlers.NewImplementation(s)

	r := chi.NewRouter()

	r.Get("/ping", handlers.PingHandler)
	r.Post("/api/tenders/new", impl.CreateTenderHandler)

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
