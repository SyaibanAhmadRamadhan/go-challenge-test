package usecase

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"

	"challenge-test-synapsis/conf"
)

type Jwt struct {
	ID   string
	Key  string
	Exp  time.Duration
	Conf conf.JwtConf
}

type JwtRTATResult struct {
	RT    string
	AT    string
	RTexp time.Duration
	ATexp time.Duration
}

func (j *Jwt) ATDefault(id string) *Jwt {
	return &Jwt{
		ID:  id,
		Key: j.Conf.ATkey,
		Exp: j.Conf.ATexp,
	}
}

func (j *Jwt) RTDefault(id string, rememberMe bool) *Jwt {
	var exp time.Duration
	if rememberMe {
		exp = j.Conf.RememberMeExp
	} else {
		exp = j.Conf.RTexp
	}

	return &Jwt{
		ID:  id,
		Key: j.Conf.RTkey,
		Exp: exp,
	}
}

func GenerateJwtHS256(jwtModel *Jwt) (str string, err error) {
	timeNow := time.Now()
	timeExp := timeNow.Add(jwtModel.Exp).Unix()

	tokenParse := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": timeExp,
		"sub": jwtModel.ID,
	})

	str, err = tokenParse.SignedString([]byte(jwtModel.Key))
	if err != nil {
		log.Warn().Msgf("failed signing Token string hs 256 | err: %v", err)
		return
	}

	return
}

func GenerateRTAT(id string, rememberMe bool) (res JwtRTATResult, err error) {
	var jwtModel *Jwt

	atModel := jwtModel.ATDefault(id)
	at, err := GenerateJwtHS256(atModel)
	if err != nil {
		return
	}

	rtModel := jwtModel.RTDefault(id, rememberMe)
	rt, err := GenerateJwtHS256(rtModel)
	if err != nil {
		return
	}

	res = JwtRTATResult{
		RT:    rt,
		AT:    at,
		RTexp: rtModel.Exp,
		ATexp: atModel.Exp,
	}

	return
}

func ClaimJwtHS256(tokenStr string, key string) (res map[string]any, err error) {
	res = map[string]any{}

	tokenParse, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Warn().Msgf("unexpected signing method : %v", token.Header["alg"])
			err = ErrJwtSigningMethodInvalid
			return nil, err
		}

		return []byte(key), nil
	})

	if err != nil {
		log.Info().Msgf("failed parse token | err: %v", err)
		return
	}

	res, _ = tokenParse.Claims.(jwt.MapClaims)
	return
}
