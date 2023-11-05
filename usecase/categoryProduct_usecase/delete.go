package categoryProduct_usecase

import (
	"context"

	"github.com/jackc/pgx/v5"

	"challenge-test-synapsis/repository"
	"challenge-test-synapsis/usecase"
)

func (c *CategoryProductUsecaseImpl) Delete(ctx context.Context, id string, param *usecase.CommonParam) (err error) {
	filters := &[]repository.Filter{
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

	err = c.categoryProductRepo.StartTx(ctx, pgx.TxOptions{
		IsoLevel:   pgx.ReadCommitted,
		AccessMode: pgx.ReadWrite,
	}, func() error {
		exist, err := c.categoryProductRepo.CheckOne(ctx, filters)
		if err != nil {
			return err
		}
		if !exist {
			return usecase.ErrCategoryProductNotFound
		}

		err = c.categoryProductRepo.Delete(ctx, id, param.UserID)

		return err
	})

	return
}
