package redis_repository

import (
	"context"

	"github.com/rs/zerolog/log"
)

func (r *RedisRepositoryImpl) Get(ctx context.Context, key string) (val string, err error) {
	val, err = r.rc.Get(ctx, key).Result()
	if err != nil {
		log.Info().Msgf("failed get in redis | err: %v", err)
	}

	return
}
