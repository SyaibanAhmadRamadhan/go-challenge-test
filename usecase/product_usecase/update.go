package product_usecase

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"

	"challenge-test-synapsis/helper"
	"challenge-test-synapsis/repository"
	"challenge-test-synapsis/usecase"
)

func (c *ProductUsecaseImpl) Update(ctx context.Context, id string, param *usecase.ProductParam) (res *usecase.ProductResult, err error) {
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
			Column:   "category_product_id",
			Value:    param.CategoryProductID,
			Operator: repository.Equality,
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
		exist, err := c.productRepo.CheckOne(ctx, filtersCheckData)
		if err != nil {
			return err
		}
		if !exist {
			return usecase.ErrProductNotFound
		}

		exist, err = c.productRepo.CheckOne(ctx, filtersCheckName)
		if err != nil {
			return err
		}
		if exist {
			return usecase.ErrProductNameIsExist
		}

		err = c.productRepo.Update(ctx, &repository.Product{
			ID:          id,
			Name:        param.Name,
			Price:       param.Price,
			Description: param.Description,
			Audit: repository.Audit{
				UpdatedAt: time.Now().Unix(),
				UpdatedBy: helper.NewNullString(param.UserID),
			},
		})

		return err
	})

	res = &usecase.ProductResult{
		ID:                id,
		Name:              param.Name,
		Price:             param.Price,
		Description:       param.Description,
		CategoryProductID: param.CategoryProductID,
	}
	return
}
