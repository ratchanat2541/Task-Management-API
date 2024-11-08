package main

import (
	"task-management-api/internal/app/router"
	"task-management-api/internal/config"
	"task-management-api/internal/dbsql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// @title Task Management API
// @description This is a description of Task Management API
// @version 1.0
// @host localhost:5555
// @BasePath /api/v1
func main() {
	var err error

	config.InitConfig()

	err = dbsql.InitDB()
	if err != nil {
		panic(err)
	}

	app := NewFiberApp()

	router.SetupRoutes(app)

	err = app.Listen(":5555")
	panic(err)
}

func NewFiberApp() *fiber.App {
	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		c.Set("X-Content-Type-Options", "nosniff")
		return c.Next()

	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",                            // specify the allowed origins
		AllowMethods: "GET, POST, PUT, DELETE",       // specify the allowed HTTP methods
		AllowHeaders: "Origin, Content-Type, Accept", // specify the allowed headers
	}))

	return app
}
