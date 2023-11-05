package categoryProduct_usecase

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"

	"challenge-test-synapsis/helper"
	"challenge-test-synapsis/repository"
	"challenge-test-synapsis/usecase"
)

func (c *CategoryProductUsecaseImpl) Create(ctx context.Context, param *usecase.CategoryProductParam) (res *usecase.CategoryProductResult, err error) {
	filters := &[]repository.Filter{
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

	id, _ := helper.NewUlid("")

	err = c.categoryProductRepo.StartTx(ctx, pgx.TxOptions{
		IsoLevel:   pgx.ReadCommitted,
		AccessMode: pgx.ReadWrite,
	}, func() error {
		exist, err := c.categoryProductRepo.CheckOne(ctx, filters)
		if err != nil {
			return err
		}
		if exist {
			return usecase.ErrCategoryProductNameIsExist
		}

		err = c.categoryProductRepo.Create(ctx, &repository.CategoryProduct{
			ID:   id,
			Name: param.Name,
			Audit: repository.Audit{
				CreatedAt: time.Now().Unix(),
				CreatedBy: param.UserID,
				UpdatedAt: time.Now().Unix(),
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
