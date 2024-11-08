package router

import (
	"task-management-api/internal/app/handler"

	"github.com/gofiber/fiber/v2"

	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func SetupRoutes(app *fiber.App) {

	// handle
	taskHandler := handler.NewTaskHandler()

	// Serve Swagger UI at /swagger/* path
	app.Get("/swagger/*", fiberSwagger.WrapHandler)
	// app.Get("/swagger/*", func(c *fiber.Ctx) error {
	// 	return c.SendFile("./doc/swagger.json")
	// })

	api := app.Group("api")

	v1 := api.Group("v1")

	// Admin API
	taskGroup := v1.Group("task", JWTAuthMiddleware)
	{
		taskGroup.Get("/", taskHandler.ListTasks)
		taskGroup.Get("/:id", taskHandler.GetTaskByID)
		taskGroup.Post("/", taskHandler.CreateTask)
		taskGroup.Put("/:id", taskHandler.UpdateTask)
		taskGroup.Put("/:id/status", taskHandler.UpdateTaskStatus)
		taskGroup.Delete("/:id", taskHandler.DeleteTaskByID)
	}
}
