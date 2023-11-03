package repository

type CategoryProduct struct {
	ID   string `sql:"id"`
	Name string `sql:"name"`
	Audit
}

type CategoryProductRepository interface {
}
