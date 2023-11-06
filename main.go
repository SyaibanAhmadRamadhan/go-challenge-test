package main

import (
	"errors"
	"os"
	"os/signal"

	"github.com/rs/zerolog/log"

	"challenge-test-synapsis/conf"
	"challenge-test-synapsis/helper"
	"challenge-test-synapsis/infra"
	"challenge-test-synapsis/presentation/rapi"
	"challenge-test-synapsis/repository/cart_repository"
	"challenge-test-synapsis/repository/categoryProduct_repository"
	"challenge-test-synapsis/repository/product_repository"
	"challenge-test-synapsis/repository/redis_repository"
	"challenge-test-synapsis/repository/session_repository"
	"challenge-test-synapsis/repository/uow_repository"
	"challenge-test-synapsis/repository/user_repository"
	"challenge-test-synapsis/usecase/auth_usecase"
	"challenge-test-synapsis/usecase/cart_usecase"
	"challenge-test-synapsis/usecase/categoryProduct_usecase"
	"challenge-test-synapsis/usecase/product_usecase"
)

func main() {
	args := os.Args
	if len(args) < 2 || args[1] != "web" && args[1] != "migrate" {
		helper.PanicIf(errors.New("must go run main.go with web or migrate | example: go run main.go migrate"))
	}

	conf.InitLogger()
	conf.LoadEnv()

	if args[1] == "migrate" {
		if len(args) == 3 {
			infra.MigrateMaster(args[2], "", "")
		} else if len(args) == 4 {
			infra.MigrateMaster(args[2], args[3], "")
		} else {
			infra.MigrateMaster("", "", "")
		}
		return
	}

	pgConn := infra.OpenConnectionDB(conf.EnvPostgresConf())
	defer func() {
		pgConn.Close()
	}()

	rcConn := infra.OpenConnectionRedis(conf.EnvRedisConf())
	defer func() {
		err := rcConn.Close()
		if err != nil {
			log.Warn().Msgf("failed closed redis connection | err: %v", err)
		}
	}()

	uow := uow_repository.NewUnitOfWorkRepositoryImpl(pgConn)
	redisRepo := redis_repository.NewRedisRepositoryImpl(rcConn)
	userRepo := user_repository.NewUserRepositoryImpl(uow)
	sessionRepo := session_repository.NewSessionRepositoryImpl(uow)
	categoryProductRepo := categoryProduct_repository.NewCategoryProductRepositoryImpl(uow)
	productRepo := product_repository.NewProductRepositoryImpl(uow)
	cartRepo := cart_repository.NewCartRepositoryImpl(uow)

	authUsecase := auth_usecase.NewAuthUsecaseImpl(userRepo, sessionRepo, redisRepo)
	categoryProductUsecase := categoryProduct_usecase.NewCategoryProductUsecaseImpl(categoryProductRepo, productRepo)
	productUsecase := product_usecase.NewProductUsecaseImpl(productRepo, categoryProductRepo)
	cartUsecase := cart_usecase.NewCartUsecaseImpl(cartRepo, productRepo)

	presenter := rapi.Presenter{
		AuthUsecase:            authUsecase,
		CategoryProductUsecase: categoryProductUsecase,
		ProductUsecase:         productUsecase,
		CartUsecase:            cartUsecase,
	}
	presenterConfig := rapi.PresenterConfig{
		WebConf:   conf.EnvWebConf(),
		Presenter: &presenter,
	}
	app := rapi.NewPresenter(presenterConfig)

	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, os.Interrupt)
	go func() {
		<-exitSignal
		log.Info().Msgf("Interrupt signal recivied, existing...")

		if err := app.Shutdown(); err != nil {
			log.Err(err).Msgf("failed gracefull shutdown")
		}
	}()

	err := app.Listen(presenterConfig.WebConf.ListenerAddr())
	helper.PanicIf(err)
}
