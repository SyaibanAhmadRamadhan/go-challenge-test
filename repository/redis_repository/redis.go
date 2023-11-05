package redis_repository

import (
	"github.com/redis/go-redis/v9"

	"challenge-test-synapsis/repository"
)

type RedisRepositoryImpl struct {
	rc *redis.Client
}

func NewRedisRepositoryImpl(rc *redis.Client) repository.RedisRepository {
	return &RedisRepositoryImpl{
		rc: rc,
	}
}
