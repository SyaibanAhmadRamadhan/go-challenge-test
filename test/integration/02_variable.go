package integration

import (
	"database/sql"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"challenge-test-synapsis/repository"
)

var db *pgxpool.Pool
var UOW repository.UOWRepository
var UserRepository repository.UserRepository
var SessionRepository repository.SessionRepository
var CategoryProductRepository repository.CategoryProductRepository

var timeUnix = time.Now().Unix()

var auditDefault = repository.Audit{
	CreatedAt: timeUnix,
	CreatedBy: "",
	UpdatedAt: timeUnix,
	UpdatedBy: sql.NullString{},
	DeletedAt: sql.NullInt64{},
	DeletedBy: sql.NullString{},
}
