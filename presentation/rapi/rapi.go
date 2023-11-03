package rapi

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"challenge-test-synapsis/conf"
)

type Presenter struct {
}

type PresenterConfig struct {
	WebConf   conf.WebConf
	Presenter *Presenter
}

func NewPresenter(config PresenterConfig) *fiber.App {
	app := fiber.New()
	fiberConfig(app)

	app.Get("/", func(ctx *fiber.Ctx) error {
		time.Sleep(10 * time.Second)
		return ctx.JSON("hello")
	})

	return app
}

func fiberConfig(app *fiber.App) {
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, id",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowCredentials: true,
	}))
}
