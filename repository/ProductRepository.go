package repository

type Product struct {
	ID          string `sql:"id"`
	Name        string `sql:"name"`
	Price       int    `sql:"price"`
	Description string `sql:"description"`
	Audit
	CategoryProduct *CategoryProduct
}

type ProductRepository interface {
	UOWRepository
}
