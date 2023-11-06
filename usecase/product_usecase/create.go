package product_usecase

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"

	"challenge-test-synapsis/helper"
	"challenge-test-synapsis/repository"
	"challenge-test-synapsis/usecase"
)

func (c *ProductUsecaseImpl) Create(ctx context.Context, param *usecase.ProductParam) (res *usecase.ProductResult, err error) {
	filters := &[]repository.Filter{
		{
			Column:   "name",
			Value:    param.Name,
			Operator: repository.Equality,
		},
		{
			Column:   "category_product_id",
			Value:    param.CategoryProductID,
			Operator: repository.Equality,
		},
		{
			Column:   "deleted_at",
			Operator: repository.IsNULL,
		},
	}
	filtersCategoryProductID := &[]repository.Filter{
		{
			Column:   "id",
			Value:    param.CategoryProductID,
			Operator: repository.Equality,
		},
		{
			Column:   "deleted_at",
			Operator: repository.IsNULL,
		},
	}

	id, _ := helper.NewUlid("")

	err = c.productRepo.StartTx(ctx, pgx.TxOptions{
		IsoLevel:   pgx.ReadCommitted,
		AccessMode: pgx.ReadWrite,
	}, func() error {
		exist, err := c.categoryProductRepo.CheckOne(ctx, filtersCategoryProductID)
		if err != nil {
			return err
		}
		if !exist {
			return usecase.ErrCategoryProductNotFound
		}

		exist, err = c.productRepo.CheckOne(ctx, filters)
		if err != nil {
			return err
		}
		if exist {
			return usecase.ErrProductNameIsExist
		}

		err = c.productRepo.Create(ctx, &repository.Product{
			ID:                id,
			CategoryProductID: param.CategoryProductID,
			Name:              param.Name,
			Stock:             param.Stock,
			Price:             param.Price,
			Description:       param.Description,
			Audit: repository.Audit{
				CreatedAt: time.Now().Unix(),
				CreatedBy: param.UserID,
				UpdatedAt: time.Now().Unix(),
			},
		})

		return err
	})

	res = &usecase.ProductResult{
		ID:                id,
		Stock:             param.Stock,
		Name:              param.Name,
		Price:             param.Price,
		Description:       param.Description,
		CategoryProductID: param.CategoryProductID,
	}
	return
}
