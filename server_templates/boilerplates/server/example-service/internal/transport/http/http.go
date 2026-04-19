package http

import (
	"example-service/internal/services"
	"example-service/internal/transport/http/controllers"

	"github.com/gofiber/fiber/v2"
)

type HTTP struct {
	app         *fiber.App
	controllers *controllers.Controllers
	//authMiddleware          func(c *fiber.Ctx) error
}

func Init(services *services.Services, app *fiber.App /*authMiddleware func(c *fiber.Ctx) error*/) *HTTP {
	return &HTTP{
		app:         app,
		controllers: controllers.Init(),
	}
}

func (http *HTTP) Start() {

}
