package repository

import (
	"database/sql"
	"time"

	"challenge-test-synapsis/repository"
)

var TimeUnix = time.Now().Unix()

var AuditDefault = repository.Audit{
	CreatedAt: TimeUnix,
	CreatedBy: "",
	UpdatedAt: TimeUnix,
	UpdatedBy: sql.NullString{},
	DeletedAt: sql.NullInt64{},
	DeletedBy: sql.NullString{},
}
