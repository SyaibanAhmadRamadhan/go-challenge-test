package cart_usecase

import (
	"context"
	"strconv"

	"github.com/jackc/pgx/v5"

	"challenge-test-synapsis/repository"
	"challenge-test-synapsis/usecase"
)

func (c *CartUsecaseImpl) DeleteItemCart(ctx context.Context, id int64, param *usecase.CommonParam) (err error) {
	filtersCheckTcart := &[]repository.Filter{
		{
			Column:   "id",
			Value:    strconv.Itoa(int(id)),
			Operator: repository.Equality,
		},
		{
			Column:   "deleted_at",
			Operator: repository.IsNULL,
		},
	}

	err = c.cartRepo.StartTx(ctx, pgx.TxOptions{
		IsoLevel:   pgx.ReadCommitted,
		AccessMode: pgx.ReadWrite,
	}, func() error {
		exist, err := c.cartRepo.CheckOneItemCart(ctx, filtersCheckTcart)
		if err != nil {
			return err
		}
		if !exist {
			return usecase.ErrItemCartNotFound
		}

		err = c.cartRepo.DeleteItemCart(ctx, id, param.UserID)
		return err
	})

	return
}
