package product_usecase

import (
	"context"

	"github.com/jackc/pgx/v5"

	"challenge-test-synapsis/repository"
	"challenge-test-synapsis/usecase"
)

func (c *ProductUsecaseImpl) Delete(ctx context.Context, id string, param *usecase.CommonParam) (err error) {
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

	err = c.productRepo.StartTx(ctx, pgx.TxOptions{
		IsoLevel:   pgx.ReadCommitted,
		AccessMode: pgx.ReadWrite,
	}, func() error {
		exist, err := c.productRepo.CheckOne(ctx, filters)
		if err != nil {
			return err
		}
		if !exist {
			return usecase.ErrProductNotFound
		}

		err = c.productRepo.Delete(ctx, id, param.UserID)

		return err
	})

	return
}
