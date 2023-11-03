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
)

func main() {
	args := os.Args
	if len(args) < 2 || args[1] != "web" && args[1] != "migrate" {
		helper.PanicIf(errors.New("must go run main.go with web or migrate | example: go run main.go migrate"))
	}

	conf.InitLogger()
	conf.LoadEnv()
	pgConf := conf.EnvPostgresConf()

	if args[1] == "migrate" {
		infra.MigrateMaster("", "", "")
		return
	}

	_ = infra.OpenConnectionDB(pgConf)

	presenter := rapi.Presenter{}
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
