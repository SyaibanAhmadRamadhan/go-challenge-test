package rapi

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"challenge-test-synapsis/conf"
	"challenge-test-synapsis/usecase"
)

type Presenter struct {
	AuthUsecase            usecase.AuthUsecase
	CategoryProductUsecase usecase.CategoryProductUsecase
	ProductUsecase         usecase.ProductUsecase
}

type PresenterConfig struct {
	WebConf   conf.WebConf
	Presenter *Presenter
}

func NewPresenter(config PresenterConfig) *fiber.App {
	app := fiber.New()
	fiberConfig(app)

	app.Post("/auth/login", config.Presenter.Login)
	app.Post("/auth/register", config.Presenter.Register)

	app.Use(config.Presenter.Otorisasi)
	app.Get("/test", func(ctx *fiber.Ctx) error {
		return ctx.JSON(ctx.Locals("role"))
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
