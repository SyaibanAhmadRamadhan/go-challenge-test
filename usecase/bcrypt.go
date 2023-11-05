package usecase

import (
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func HashBcrypt(password string) (str string, err error) {
	res, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Info().Msgf("failed generate password hash | err:%v", err)
		return
	}

	str = string(res)

	return
}

func CompareBcrypt(hash string, password string) (b bool) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
