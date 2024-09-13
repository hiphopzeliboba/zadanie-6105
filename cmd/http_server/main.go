package main

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"mini_url/internal/config"
	tender2 "mini_url/internal/service/tender"

	"log"
	"mini_url/internal/api/handlers"
	"mini_url/internal/repository/tender"
	"net/http"
	"os"
)

func main() {
	// Считываем переменные окружения
	//pgPool, err := tender.NewPool(os.Getenv("POSTGRES_CONN"))
	//if err != nil {
	//	log.Fatalf("Unable to create database pool: %v\n", err)
	//}
	//defer pgPool.Close()
	ctx := context.Background()

	config.Load(".env")

	pool, err := pgxpool.Connect(ctx, os.Getenv("POSTGRES_CONN"))
	if err != nil {
		log.Fatal(err)
	}
	tRepo := tender.NewRepository(pool)

	s := tender2.NewService(tRepo)

	impl := handlers.NewImplementation(s)

	//pool, err := tender.NewRepository(pgPool)
	//if err != nil {
	//	log.Fatalf("Unable to create database pool: %v\n", err)
	//}

	r := chi.NewRouter()

	r.Get("/ping", handlers.PingHandler)
	r.Post("/api/tenders/new", impl.CreateTenderHandler)

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
