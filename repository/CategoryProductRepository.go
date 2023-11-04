package repository

import (
	"context"
)

type CategoryProduct struct {
	ID   string `sql:"id"`
	Name string `sql:"name"`
	Audit
} // 128 bytes mem. check for unit test TestStructEmbedMemoryUsage

type CategoryProductJoin struct {
	ID   string `sql:"id"`
	Name string `sql:"name"`
	Audit
	Product *[]Product
} // 136 bytes mem. check for unit test TestStructEmbedMemoryUsage

type CategoryProductRepository interface {
	Create(ctx context.Context, categoryProduct *CategoryProduct) (err error)
	Update(ctx context.Context, categoryProduct *CategoryProduct) (err error)
	Delete(ctx context.Context, id string, userID string) (err error)
	CheckOne(ctx context.Context, filters *[]Filter) (b bool, err error)
	FindOne(ctx context.Context, filters *[]Filter) (categoryProduct *CategoryProduct, err error)
	FindAllAndSearch(ctx context.Context, param FindAllAndSearchParam) (categoryProducts *[]CategoryProduct, total int, err error)
	UOWRepository
}
