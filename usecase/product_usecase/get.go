package product_usecase

import (
	"context"
	"math"

	"github.com/jackc/pgx/v5"

	"challenge-test-synapsis/repository"
	"challenge-test-synapsis/usecase"
)

func (c *ProductUsecaseImpl) GetAndSearch(ctx context.Context, search string, param *usecase.GetProductParam) (res *[]usecase.ProductResult, page *usecase.PaginateResult, err error) {
	filters := &[]repository.Filter{
		{
			Column:   "category_product_id",
			Value:    param.CategoryProductID,
			Operator: repository.IsNULL,
		},
		{
			Column:   "deleted_at",
			Operator: repository.IsNULL,
		},
	}

	var products *[]repository.Product

	findAllAndSearch := repository.FindAllAndSearchParam{
		Filters: filters,
		Pagination: repository.Pagination{
			Limit:  param.PageSize,
			Offset: (param.Page - 1) * param.PageSize,
			Orders: map[string]string{
				"id": "DESC",
			},
		},
		Search: search,
	}

	err = c.productRepo.StartTx(ctx, pgx.TxOptions{
		IsoLevel:   pgx.ReadCommitted,
		AccessMode: pgx.ReadWrite,
	}, func() error {
		data, total, err := c.productRepo.FindAllAndSearch(ctx, findAllAndSearch)

		page = &usecase.PaginateResult{}
		page.CurrentPage = param.Page
		page.Total = total
		page.PageTotal = int(math.Ceil(float64(page.Total) / float64(param.PageSize)))
		page.PageSize = param.PageSize

		products = data
		return err
	})

	res = &[]usecase.ProductResult{}

	for _, product := range *products {
		*res = append(*res, usecase.ProductResult{
			ID:                product.ID,
			Name:              product.Name,
			Price:             product.Price,
			Description:       product.Description,
			CategoryProductID: param.CategoryProductID,
		})
	}
	return
}
