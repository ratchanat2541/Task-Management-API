package handler

import (
	"net/http"
	"task-management-api/internal/app/model"
	"task-management-api/internal/app/service"

	"github.com/gofiber/fiber/v2"
)

type TaskHandler struct {
	taskService *service.TaskService
}

func NewTaskHandler() *TaskHandler {
	return &TaskHandler{
		taskService: service.NewTaskService(),
	}
}

// ListTasks returns tasks
// @Summary Lust tasks
// @Description Lust tasks
// @Tags task
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Success 200 {object} model.Response{result=model.Response} "Success"
// @Failure 500 {object} model.Response{error=model.Response} "Internal Server Error"
// @Router /task [get]
func (h *TaskHandler) ListTasks(c *fiber.Ctx) error {
	tasks, err := h.taskService.ListTasks()
	if err != nil {
		return c.JSON(model.Response{Code: http.StatusInternalServerError, Message: "Error Get all", Data: err.Error()})
	}

	countAll := int64(len(tasks))
	return c.JSON(model.Response{Code: http.StatusOK, Data: tasks, CountAll: &countAll})
}

// GetTaskByID returns task
// @Summary Get one task
// @Description Get one task
// @Tags task
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param id path string true "task id"
// @Success 200 {object} model.Response{result=model.Response} "Success"
// @Failure 400 {object} model.Response{error=model.Response} "Bad Request"
// @Failure 500 {object} model.Response{error=model.Response} "Internal Server Error"
// @Router /task/{id} [get]
func (h *TaskHandler) GetTaskByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.JSON(model.Response{Code: http.StatusBadRequest, Message: "Error ParamsInt", Data: err.Error()})
	}

	admin, err := h.taskService.GetTaskByID(id)
	if err != nil {
		return c.JSON(model.Response{Code: http.StatusInternalServerError, Message: "Error Get by id", Data: err.Error()})
	}

	return c.JSON(model.Response{Code: http.StatusOK, Data: admin})
}

// CreateTask
// @Summary Create task
// @Description Create task
// @Tags task
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param taskRequest body model.TaskRequest true "Create task"
// @Success 200 {object} model.Response{result=model.Response} "Success"
// @Failure 400 {object} model.Response{error=model.Response} "Bad Request"
// @Failure 500 {object} model.Response{error=model.Response} "Internal Server Error"
// @Router /task [post]
func (h *TaskHandler) CreateTask(c *fiber.Ctx) error {
	taskRequest := new(model.TaskRequest)
	err := c.BodyParser(taskRequest)
	if err != nil {
		return c.JSON(model.Response{Code: http.StatusBadRequest, Message: "Error BodyParser", Data: err.Error()})
	}

	task, err := h.taskService.CreateTask(*taskRequest)
	if err != nil {
		return c.JSON(model.Response{Code: http.StatusInternalServerError, Message: "Error Update status", Data: err.Error()})
	}

	return c.JSON(model.Response{Code: http.StatusOK, Message: "Create success", Data: task})
}

// UpdateTask
// @Summary Update task
// @Description Update task
// @Tags task
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param id path string true "user uuid"
// @Param taskUpdateRequest body model.TaskUpdateRequest true "Update task"
// @Success 200 {object} model.Response{result=model.Response} "Success"
// @Failure 400 {object} model.Response{error=model.Response} "Bad Request"
// @Failure 500 {object} model.Response{error=model.Response} "Internal Server Error"
// @Router /task/{id} [put]
func (h *TaskHandler) UpdateTask(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.JSON(model.Response{Code: http.StatusBadRequest, Message: "Error ParamsInt", Data: err.Error()})
	}

	taskUpdateRequest := new(model.TaskUpdateRequest)
	err = c.BodyParser(taskUpdateRequest)
	if err != nil {
		return c.JSON(model.Response{Code: http.StatusBadRequest, Message: "Error BodyParser", Data: err.Error()})
	}

	err = h.taskService.UpdateTask(id, *taskUpdateRequest)
	if err != nil {
		return c.JSON(model.Response{Code: http.StatusInternalServerError, Message: "Error Update status", Data: err.Error()})
	}

	return c.JSON(model.Response{Code: http.StatusOK, Message: "Update success"})
}

// UpdateTaskStatus
// @Summary Update status
// @Description Update status
// @Tags task
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param id path string true "user uuid"
// @Param taskUpdateStatusRequest body model.TaskUpdateStatusRequest true "Update User status"
// @Success 200 {object} model.Response{result=model.Response} "Success"
// @Failure 400 {object} model.Response{error=model.Response} "Bad Request"
// @Failure 500 {object} model.Response{error=model.Response} "Internal Server Error"
// @Router /task/{id}/status [put]
func (h *TaskHandler) UpdateTaskStatus(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.JSON(model.Response{Code: http.StatusBadRequest, Message: "Error ParamsInt", Data: err.Error()})
	}

	taskUpdateStatusRequest := new(model.TaskUpdateStatusRequest)
	err = c.BodyParser(taskUpdateStatusRequest)
	if err != nil {
		return c.JSON(model.Response{Code: http.StatusBadRequest, Message: "Error BodyParser", Data: err.Error()})
	}

	err = h.taskService.UpdateTaskStatus(id, taskUpdateStatusRequest.Status)
	if err != nil {
		return c.JSON(model.Response{Code: http.StatusInternalServerError, Message: "Error Update status", Data: err.Error()})
	}

	return c.JSON(model.Response{Code: http.StatusOK, Message: "Update success"})
}

// DeleteTaskByID
// @Summary Delete task by id
// @Description Delete task by id
// @Tags task
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param id path string true "task id"
// @Success 200 {object} model.Response{result=model.Response} "Success"
// @Failure 400 {object} model.Response{error=model.Response} "Bad Request"
// @Failure 500 {object} model.Response{error=model.Response} "Internal Server Error"
// @Router /task/{id} [delete]
func (h *TaskHandler) DeleteTaskByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.JSON(model.Response{Code: http.StatusBadRequest, Message: "Error ParamsInt", Data: err.Error()})
	}

	err = h.taskService.DeleteTaskByID(id)
	if err != nil {
		return c.JSON(model.Response{Code: http.StatusInternalServerError, Message: "Error Delete", Data: err.Error()})
	}

	return c.JSON(model.Response{Code: http.StatusOK, Data: "Delete success"})
}
