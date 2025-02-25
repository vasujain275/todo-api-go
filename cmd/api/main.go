package main

import (
	"github.com/joho/godotenv"
	"github.com/vasujain275/todo-api-go/internal/env"
	"log"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("error loading env file: %v", err)
	}

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:            env.GetString("DB_ADDR", "postgres://postgres:password@localhost:5432/todo?sslmode=disable"),
			maxOpenConns:    env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns:    env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxConnLifetime: env.GetString("DB_MAX_CONN_LIFETIME", "1h"),
		},
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
