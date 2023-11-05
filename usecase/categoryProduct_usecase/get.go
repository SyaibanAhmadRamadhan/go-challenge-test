package categoryProduct_usecase

import (
	"context"
	"math"

	"github.com/jackc/pgx/v5"

	"challenge-test-synapsis/repository"
	"challenge-test-synapsis/usecase"
)

func (c *CategoryProductUsecaseImpl) GetAndSearch(ctx context.Context, search string, param *usecase.PaginateParam) (res *[]usecase.CategoryProductResult, page *usecase.PaginateResult, err error) {
	filters := &[]repository.Filter{
		{
			Column:   "deleted_at",
			Operator: repository.IsNULL,
		},
	}

	var categoryProduts *[]repository.CategoryProduct

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

	err = c.categoryProductRepo.StartTx(ctx, pgx.TxOptions{
		IsoLevel:   pgx.ReadCommitted,
		AccessMode: pgx.ReadWrite,
	}, func() error {
		data, total, err := c.categoryProductRepo.FindAllAndSearch(ctx, findAllAndSearch)

		page = &usecase.PaginateResult{}
		page.CurrentPage = param.Page
		page.Total = total
		page.PageTotal = int(math.Ceil(float64(page.Total) / float64(param.PageSize)))
		page.PageSize = param.PageSize

		categoryProduts = data
		return err
	})

	res = &[]usecase.CategoryProductResult{}

	for _, categoryProdut := range *categoryProduts {
		*res = append(*res, usecase.CategoryProductResult{
			ID:   categoryProdut.ID,
			Name: categoryProdut.Name,
		})
	}
	return
}
