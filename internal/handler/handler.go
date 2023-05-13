package handler

import (
	"github.com/gin-gonic/gin"
	"todo-app/internal/service"
)

type Handler struct {
	services *service.Service
}

func New(s *service.Service) *Handler {
	return &Handler{services: s}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())

	api := router.Group("/api")

	auth := api.Group("/auth")
	{
		auth.POST("/sing-in", h.singIn)
		auth.POST("/sing-up", h.singUp)
	}

	tasks := api.Group("/tasks", h.userIdentity)
	{
		tasks.GET("/", h.getAllTask)
		tasks.POST("/", h.saveTask)
		tasks.GET("/:id", h.getTask)
		tasks.DELETE("/:id", h.deleteTask)
		tasks.PUT("/:id", h.updateTask)
	}

	return router
}
