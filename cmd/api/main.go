package main

import (
	"Social/internal/db"
	"Social/internal/env"
	"Social/internal/store"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, using defaults/system env")
	}

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:               env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable"),
			maxOpenConnections: env.GetInt("DB_MAX_OPEN_CONNECTIONS", 30),
			maxIdleConnections: env.GetInt("DB_MAX_IDLE_CONNECTIONS", 30),
			maxIdleTime:        env.GetString("DB_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConnections,
		cfg.db.maxIdleConnections,
		cfg.db.maxIdleTime,
	)

	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	log.Println("databse Pool established")

	store := store.NewPostgresStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
