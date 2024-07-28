package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/swaggo/http-swagger"
	"log"
	"os"
	"test_posiflora/app/listener"
	"test_posiflora/app/middleware"
	"test_posiflora/config"
	"test_posiflora/config/router"
	_ "test_posiflora/docs"
)

// @title                       test Posiflora
// @version                     1.0.0
// @in                          header
// @name                        Authorization
// @BasePath                    /
// @securityDefinitions.basic   ApiKeyAuth
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("Not found .env", err)
	}
	app := fiber.New(config.FiberConfig())
	middleware.AppMiddleware(app)
	router.AppRouter(app)

	config.AppConfig()
	go listener.ItemTaskListener(config.RabbitMQ)
	err := app.Listen(fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT")))

	if err != nil {
		log.Panicf("Server is not running: %v", err.Error())
	}
	fmt.Println("Hello World")
}
