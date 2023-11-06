package conf

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func LoadEnv() {
	dirList := []string{"", "../", "../../", "../../../"}
	for _, dir := range dirList {
		envFile := fmt.Sprintf("%s.env", dir)
		err := godotenv.Load(envFile)
		if err == nil {
			// envFileOverride := fmt.Sprintf("%s.env.override", dir)
			// err = godotenv.Load(envFileOverride)
			// helper.PanicIf(err)

			log.Info().Msg("initialization .env successfully")
			return
		}
	}

	panic("cannot load file .env and .env.override")
}
