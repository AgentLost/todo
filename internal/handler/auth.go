package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strings"
	"todo-app/internal/model"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(h.services.AuthService.GetHeader())
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}
	log.Println("........................")

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	log.Println("........................")

	userId, err := h.services.AuthService.GetUserId(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	log.Printf("token = %s, userId = %d", headerParts[1], userId)
	c.Set("userId", userId)
}

func (h *Handler) singIn(c *gin.Context) {
	input := model.AuthRequest{}

	if err := c.ShouldBindBodyWith(&input, binding.JSON); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.UserService.GetByEmail(input.Email)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid password or email")
		return
	}

	token, err := h.services.AuthService.Generate(user.Id)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response := model.AuthResponse{Token: token}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) singUp(c *gin.Context) {
	input := model.UserRequest{}

	if err := c.ShouldBindBodyWith(&input, binding.JSON); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	log.Println("..................")

	err := h.services.UserService.Register(input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "ok")
}
