package repository

import (
	"context"
	"time"
)

type RedisSet struct {
	Key   string
	Value any
	Exp   time.Duration
}

type RedisRepository interface {
	Set(ctx context.Context, set RedisSet) (err error)
	Get(ctx context.Context, key string) (val string, err error)
}
