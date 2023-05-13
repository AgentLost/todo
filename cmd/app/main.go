package main

import (
	"github.com/caarlos0/env"
	"log"
	"todo-app/config"
	"todo-app/internal/app"
)

func main() {
	cfg := &config.Config{}

	if err := env.Parse(cfg); err != nil {
		log.Fatal(err)
	}

	app.Run(cfg)
}
