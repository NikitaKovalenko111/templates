package example_controller

import (
	example_dto "example-service/internal/dto/example"
	_ "example-service/internal/types"

	"github.com/gofiber/fiber/v2"
)

type ExampleController struct {
	// Here are services associated with the controller
}

func Init() *ExampleController {
	return &ExampleController{}
}

func (controller *ExampleController) Start(route string, app *fiber.App /*authMiddleware func(c *fiber.Ctx) error*/) {
	router := app.Group("/" + route)

	router.Get("/health", controller.HealthCheck)
}

// HealthCheck godoc
//
//	@Summary		Check health
//	@Description	Checks health of the server.
//	@Tags			health
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	example_dto.HealthCheckResponse
//	@Router			/health [get]
func (controller *ExampleController) HealthCheck(c *fiber.Ctx) error {
	var response example_dto.HealthCheckResponse

	response = example_dto.HealthCheckResponse{
		Status: fiber.StatusOK,
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
