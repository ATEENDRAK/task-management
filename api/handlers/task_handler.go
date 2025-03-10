package handlers

import (
	"net/http"
	"strconv"
	"task-mgmt-service/domain"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	Service domain.TaskService
}

// NewTaskHandler creates a new TaskHandler
func NewTaskHandler(service domain.TaskService) *TaskHandler {
	return &TaskHandler{Service: service}
}

// Create a new task
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var req struct {
		Title  string `json:"title"`
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task, err := h.Service.CreateTask(req.Title, req.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create task"})
		return
	}
	c.JSON(http.StatusCreated, task)
}

// Get all tasks with pagination & filtering
func (h *TaskHandler) GetTasks(c *gin.Context) {
	status := c.Query("status")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	tasks, err := h.Service.GetTasks(status, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// Update a task
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req struct {
		Title  string `json:"title"`
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.Service.UpdateTask(uint(id), req.Title, req.Status)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}

// Delete a task
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.Service.DeleteTask(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
