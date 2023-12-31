package repository

import (
	"context"
)

type Product struct {
	ID                string `sql:"id"`
	CategoryProductID string `sql:"category_product_id"`
	Name              string `sql:"name"`
	Stock             int    `sql:"stock"`
	Price             int    `sql:"price"`
	Description       string `sql:"description"`
	Audit
}

type ProductRepository interface {
	Create(ctx context.Context, product *Product) (err error)
	Update(ctx context.Context, product *Product) (err error)
	Delete(ctx context.Context, id string, userID string) (err error)
	CheckOne(ctx context.Context, filters *[]Filter) (b bool, err error)
	FindOne(ctx context.Context, filters *[]Filter) (product *Product, err error)
	FindAllAndSearch(ctx context.Context, param FPSParam) (products *[]Product, total int, err error)
	UOWRepository
}
