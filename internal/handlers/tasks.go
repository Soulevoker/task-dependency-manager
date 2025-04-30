package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task-dependency-manager/internal/models"
	"task-dependency-manager/internal/services"
)

type TaskHandler struct {
	service *services.TaskService
}

func NewTaskHandler(service *services.TaskService) *TaskHandler {
	return &TaskHandler{
		service: service,
	}
}

func (h *TaskHandler) GetTask(c *gin.Context) {
	task, err := h.service.GetTask(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	createdTask, err := h.service.CreateTask(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdTask)
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	var task models.Task
	task.ID = c.Param("id")
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	updatedTask, err := h.service.UpdateTask(&task)
	if err != nil {
		if err.Error() == "task not found" {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}
	c.IndentedJSON(http.StatusOK, updatedTask)
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	taskID := c.Param("id")
	if err := h.service.DeleteTask(taskID); err != nil {
		if err.Error() == "task not found" {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "task {" + taskID + "} deleted"})
}

func (h *TaskHandler) ListTasks(c *gin.Context) {
	l, err := h.service.ListTasks()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, l)
}
