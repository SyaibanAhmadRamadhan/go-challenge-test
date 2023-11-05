package integration

import (
	"database/sql"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"

	"challenge-test-synapsis/repository"
	"challenge-test-synapsis/usecase"
)

var db *pgxpool.Pool
var rc *redis.Client
var UOW repository.UOWRepository
var UserRepository repository.UserRepository
var SessionRepository repository.SessionRepository
var CategoryProductRepository repository.CategoryProductRepository
var ProductRepository repository.ProductRepository
var RedisRepository repository.RedisRepository

var AuthUsecase usecase.AuthUsecase

var timeUnix = time.Now().Unix()

var auditDefault = repository.Audit{
	CreatedAt: timeUnix,
	CreatedBy: "",
	UpdatedAt: timeUnix,
	UpdatedBy: sql.NullString{},
	DeletedAt: sql.NullInt64{},
	DeletedBy: sql.NullString{},
}
