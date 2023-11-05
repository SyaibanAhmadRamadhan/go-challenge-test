package categoryProduct_usecase

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"

	"challenge-test-synapsis/helper"
	"challenge-test-synapsis/repository"
	"challenge-test-synapsis/usecase"
)

func (c *CategoryProductUsecaseImpl) Update(ctx context.Context, id string, param *usecase.CategoryProductParam) (res *usecase.CategoryProductResult, err error) {
	filtersCheckData := &[]repository.Filter{
		{
			Column:   "id",
			Value:    id,
			Operator: repository.Equality,
		},
		{
			Column:   "deleted_at",
			Operator: repository.IsNULL,
		},
	}
	filtersCheckName := &[]repository.Filter{
		{
			Column:   "id",
			Value:    id,
			Operator: repository.Inequality,
		},
		{
			Column:   "name",
			Value:    param.Name,
			Operator: repository.Equality,
		},
		{
			Column:   "deleted_at",
			Operator: repository.IsNULL,
		},
	}

	err = c.categoryProductRepo.StartTx(ctx, pgx.TxOptions{
		IsoLevel:   pgx.ReadCommitted,
		AccessMode: pgx.ReadWrite,
	}, func() error {
		exist, err := c.categoryProductRepo.CheckOne(ctx, filtersCheckData)
		if err != nil {
			return err
		}
		if !exist {
			return usecase.ErrCategoryProductNotFound
		}

		exist, err = c.categoryProductRepo.CheckOne(ctx, filtersCheckName)
		if err != nil {
			return err
		}
		if exist {
			return usecase.ErrCategoryProductNameIsExist
		}

		err = c.categoryProductRepo.Update(ctx, &repository.CategoryProduct{
			ID:   id,
			Name: param.Name,
			Audit: repository.Audit{
				UpdatedAt: time.Now().Unix(),
				UpdatedBy: helper.NewNullString(param.UserID),
			},
		})

		return err
	})

	res = &usecase.CategoryProductResult{
		ID:   id,
		Name: param.Name,
	}
	return
}
