package router

import (
	"github.com/gofiber/fiber/v2"
	"test_posiflora/app/controllers"
	"test_posiflora/app/controllers/taskController"
)

func apiRouter(f fiber.Router) {
	appController := initializationController()

	task := f.Group("/notifications")
	{
		task.Post("/", appController.TaskController.CreateOne)
		task.Get("/", appController.TaskController.GetAll)
	}
}

func initializationController() controllers.Controller {
	return controllers.Controller{
		TaskController: taskController.NewTaskController(),
	}
}
