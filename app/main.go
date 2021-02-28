package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/goku321/api-server/app/api"
	"github.com/goku321/geolocation/store"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	connStr := os.Getenv("DB_CONN_STR")
	pg, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("cannot connect to db: %s", err)
	}
	pg.SetMaxIdleConns(2)
	pg.SetMaxOpenConns(10)

	store := store.New(pg)
	defer store.Close()

	r := api.New(store)
	r.Register()

	// add ctx down the calls
	srv := &http.Server{
		Handler:      r.Root,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// Handle graceful shutdown.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block to wait for interrupt.
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	srv.Shutdown(ctx)
	log.Println("shutting down server...")
	os.Exit(0)
}
