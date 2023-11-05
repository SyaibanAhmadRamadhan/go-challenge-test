package infra

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"

	"challenge-test-synapsis/conf"
)

func OpenConnectionRedis(conf conf.RedisConf) (rc *redis.Client) {
	host := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	rc = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: conf.Pass,
		DB:       conf.DB,
	})

	ctx := context.TODO()
	ping, err := rc.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	log.Info().Msgf("connection to redis : %s", ping)
	
	return
}
