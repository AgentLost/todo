package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strconv"
	"todo-app/internal/model"
)

func (h *Handler) getTask(c *gin.Context) {
	stringId := c.Param("id")
	id, err := strconv.Atoi(stringId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "bad param")
		return
	}

	task, err := h.services.TaskService.Get(c.GetInt("userId"), id)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *Handler) getAllTask(c *gin.Context) {
	tasks, err := h.services.TaskService.GetAll(c.GetInt("userId"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "bad request")
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *Handler) deleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "bad params")
		return
	}

	err = h.services.TaskService.Delete(c.GetInt("userId"), id)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "ok")
}

func (h *Handler) saveTask(c *gin.Context) {
	var input model.TaskRequest
	if err := c.ShouldBindBodyWith(&input, binding.JSON); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	task := model.TaskDto{
		UserId:      c.GetInt("userId"),
		Title:       input.Title,
		Due:         input.Due,
		Description: input.Description,
	}

	err := h.services.TaskService.Save(task)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "ok")
}

func (h *Handler) updateTask(c *gin.Context) {
	taskId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input model.TaskRequest
	if err := c.ShouldBindBodyWith(&input, binding.JSON); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	task := model.TaskDto{
		UserId:      c.GetInt("userId"),
		Title:       input.Title,
		Due:         input.Due,
		Description: input.Description,
	}

	err = h.services.TaskService.Update(c.GetInt("userId"), taskId, task)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "ok")
}
