package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

type errResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Printf("error message %s", message)
	c.AbortWithStatusJSON(statusCode, errResponse{message})
}
