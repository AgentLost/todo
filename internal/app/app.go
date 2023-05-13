package app

import (
	"log"
	"todo-app/config"
	"todo-app/internal/handler"
	"todo-app/internal/helpers"
	"todo-app/internal/infrastructure/database"
	"todo-app/internal/infrastructure/server"
	"todo-app/internal/repo"
	"todo-app/internal/service"
)

func Run(cfg *config.Config) {
	log.Println(cfg)
	db := database.New(cfg)
	//	Migrate(cfg)

	r := repo.New(db)
	log.Printf("repo %v", r)
	p := helpers.NewTokenProvider(cfg)
	log.Printf("provider %v", p)
	s := service.New(r, p)
	log.Printf("service %v", s)
	h := handler.New(s)
	log.Printf("service %v", h)

	http := server.Server{}

	if err := http.Run(cfg.HTTPPort, h.InitRoutes()); err != nil {
		log.Fatal(err)
	}

}
