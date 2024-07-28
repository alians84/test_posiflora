package router

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func swaggerRouter(g fiber.Router) {
	g.All("*", swagger.HandlerDefault)
}
