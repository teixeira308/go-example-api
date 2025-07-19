package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/teixeira308/go-example-api/internal/store"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080" // fallback
	}

	cfg := config{
		addr: port,
	}

	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}
	log.Printf("Server at %s", app.config.addr)
	mux := app.mount()
	log.Fatal(app.run(mux))
}
