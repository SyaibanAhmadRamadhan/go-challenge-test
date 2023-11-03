package repository

type Role struct {
	ID          int    `sql:"id"`
	Name        string `sql:"name"`
	Description string `sql:"description"`
}

type RoleRepository interface {
}
