package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/vasujain275/todo-api-go/internal/db"
	"github.com/vasujain275/todo-api-go/internal/env"
	"github.com/vasujain275/todo-api-go/internal/store"
	"log"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("error loading env file: %v", err)
	}

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:            env.GetString("DB_ADDR", "postgres://postgres:postgres123@localhost:5432/todoDB?sslmode=disable"),
			maxOpenConns:    env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns:    env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxConnLifetime: env.GetString("DB_MAX_CONN_LIFETIME", "1h"),
		},
	}

	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxConnLifetime)
	defer db.Close()

	if err != nil {
		log.Panicf("error connecting to db: %v", err)
	}

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
