package app

import (
	"fmt"
	"log"
	"todo-app/config"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
)

func Migrate(cfg *config.Config) {
	log.Println("create migrate...")
	m, err := migrate.New("file://"+cfg.Dir,
		fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
			cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name, cfg.SSLMode))

	if err != nil {
		log.Println(err)
	}

	log.Println("try down migrate")
	if err := m.Down(); err != nil {
		log.Println(err)
	}

	log.Println("try up migrate")
	if err := m.Up(); err != nil {
		log.Println(err)
	}
}
