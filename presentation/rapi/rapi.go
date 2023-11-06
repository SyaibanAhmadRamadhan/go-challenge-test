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
	CartUsecase            usecase.CartUsecase
}

type PresenterConfig struct {
	WebConf   conf.WebConf
	Presenter *Presenter
}

func NewPresenter(config PresenterConfig) *fiber.App {
	app := fiber.New()
	fiberConfig(app)

	app.Post("/auth/login", CheckLogin, config.Presenter.Login)
	app.Post("/auth/register", CheckLogin, config.Presenter.Register)
	app.Get("/category-product", config.Presenter.GetCategoryProduct)
	app.Get("/product", config.Presenter.GetProduct)

	mustLogin := app.Group("", config.Presenter.Otorisasi)
	app.Get("/cart", config.Presenter.GetProductCart)
	app.Post("/cart", config.Presenter.AddProductCart)
	app.Put("/cart/:id", config.Presenter.UpdateProductCart)
	app.Delete("/cart/:id", config.Presenter.DeleteProductCart)

	mustAdmin := mustLogin.Group("", MustBeAdmin)
	mustAdmin.Post("/category-product", config.Presenter.AddCategoryProduct)
	mustAdmin.Put("/category-product/:id", config.Presenter.UpdateCategoryProduct)
	mustAdmin.Delete("/category-product/:id", config.Presenter.DeleteCategoryProduct)
	mustAdmin.Post("/product", config.Presenter.AddProduct)
	mustAdmin.Put("/product/:id", config.Presenter.UpdateProduct)
	mustAdmin.Delete("/product/:id", config.Presenter.DeleteProduct)

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
