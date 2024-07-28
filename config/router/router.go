package router

import "github.com/gofiber/fiber/v2"

func AppRouter(a *fiber.App) {
	swagger := a.Group("/swagger")
	{
		swaggerRouter(swagger)
	}
	api := a.Group("/")
	{
		apiRouter(api)
	}
}
