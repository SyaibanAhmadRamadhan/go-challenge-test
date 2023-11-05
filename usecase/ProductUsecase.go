package usecase

import (
	"context"
)

type ProductParam struct {
	CategoryProductID string
	Name              string
	Price             int
	Description       string
	CommonParam
}

type ProductResult struct {
	ID                string
	Name              string
	Price             int
	Description       string
	CategoryProductID string
}

type GetProductParam struct {
	CategoryProductID string
	PaginateParam
}

type ProductUsecase interface {
	Create(ctx context.Context, param *ProductParam) (res *ProductResult, err error)
	Update(ctx context.Context, id string, param *ProductParam) (res *ProductResult, err error)
	Delete(ctx context.Context, id string, param *CommonParam) (res *ProductResult, err error)
	GetAndSearch(ctx context.Context, search string, param *GetProductParam) (res *[]ProductResult, page *PaginateResult, err error)
}
