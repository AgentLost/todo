package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"todo-app/config"
)

type Postgres struct {
	DB *sqlx.DB
}

func New(cfg *config.Config) *Postgres {
	postgres := Postgres{}

	log.Printf("try open database cfg %v\n", cfg)
	db, err := sqlx.Open(cfg.DriverName,
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.SSLMode))

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("try ping to database=%v\n", db)
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	postgres.DB = db
	return &postgres
}
