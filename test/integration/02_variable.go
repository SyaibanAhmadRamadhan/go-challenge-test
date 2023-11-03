package integration

import (
	"database/sql"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"challenge-test-synapsis/repository"
)

var db *pgxpool.Pool
var UserRepository repository.UserRepository
var UOW repository.UOWRepository

var timeUnix = time.Now().Unix()

var auditDefault = repository.Audit{
	CreatedAt: timeUnix,
	CreatedBy: "",
	UpdatedAt: timeUnix,
	UpdatedBy: sql.NullString{},
	DeletedAt: sql.NullInt64{},
	DeletedBy: sql.NullString{},
}

var roleDefault = repository.Role{
	ID:          1,
	Name:        "",
	Description: "",
}
