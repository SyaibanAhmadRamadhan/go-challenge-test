package repository

import (
	"database/sql"
	"fmt"
)

type Audit struct {
	CreatedAt int64          `sql:"created_at"`
	CreatedBy string         `sql:"created_by"`
	UpdatedAt int64          `sql:"updated_at"`
	UpdatedBy sql.NullString `sql:"updated_by"`
	DeletedAt sql.NullInt64  `sql:"deleted_at"`
	DeletedBy sql.NullString `sql:"deleted_by"`
}

func AuditToQuery(prefix string) string {
	return fmt.Sprintf("%screated_at, %screated_by, %supdated_at, %supdated_by, %sdeleted_at, %sdeleted_by",
		prefix, prefix, prefix, prefix, prefix, prefix)
}
