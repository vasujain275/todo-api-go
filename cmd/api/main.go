package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/vasujain275/todo-api-go/internal/env"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("error loading env file: %v", err)
	}

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
