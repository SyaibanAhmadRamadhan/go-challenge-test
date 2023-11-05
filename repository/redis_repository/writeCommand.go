package redis_repository

import (
	"context"

	"github.com/rs/zerolog/log"

	"challenge-test-synapsis/repository"
)

func (r *RedisRepositoryImpl) Set(ctx context.Context, set repository.RedisSet) (err error) {
	err = r.rc.Set(ctx, set.Key, set.Value, set.Exp).Err()
	if err != nil {
		log.Warn().Msgf("failed set data to redis | err: %v", err)
	}
	
	return
}
