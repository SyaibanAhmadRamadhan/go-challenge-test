package repository

type Audit struct {
	CreatedAt int    `sql:"created_at"`
	CreatedBy string `sql:"created_by"`
	UpdatedAt int    `sql:"updated_at"`
	UpdatedBy string `sql:"updated_by"`
	DeletedAt int    `sql:"deleted_at"`
	DeletedBy string `sql:"deleted_by"`
}
