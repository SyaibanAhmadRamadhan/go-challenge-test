package usecase

import (
	"context"
)

type CategoryProductParam struct {
	Name string
	CommonParam
}

type CategoryProductResult struct {
	ID   string
	Name string
}

type CategoryProductUsecase interface {
	Create(ctx context.Context, param *CategoryProductParam) (res *CategoryProductResult, err error)
	Update(ctx context.Context, id string, param *CategoryProductParam) (res *CategoryProductResult, err error)
	Delete(ctx context.Context, id string, param *CommonParam) (err error)
	GetAndSearch(ctx context.Context, search string, param *PaginateParam) (res *[]CategoryProductResult, page *PaginateResult, err error)
}
