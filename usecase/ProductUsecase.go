package usecase

import (
	"context"
)

type ProductParam struct {
	CategoryProductID string
	Name              string
	Price             int
	Stock             int
	Description       string
	CommonParam
}

type ProductResult struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Price             int    `json:"price"`
	Stock             int    `json:"stock"`
	Description       string `json:"description"`
	CategoryProductID string `json:"category_product_id"`
}

type GetProductParam struct {
	CategoryProductID string
	PaginateParam
}

type ProductUsecase interface {
	Create(ctx context.Context, param *ProductParam) (res *ProductResult, err error)
	Update(ctx context.Context, id string, param *ProductParam) (res *ProductResult, err error)
	Delete(ctx context.Context, id string, param *CommonParam) (err error)
	GetAndSearch(ctx context.Context, search string, param *GetProductParam) (res *[]ProductResult, page *PaginateResult, err error)
}
