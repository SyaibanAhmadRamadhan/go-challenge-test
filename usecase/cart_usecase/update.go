package cart_usecase

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"

	"challenge-test-synapsis/repository"
	"challenge-test-synapsis/usecase"
)

func (c *CartUsecaseImpl) UpdateItemCart(ctx context.Context, id int64, param *usecase.ItemCartParam) (res *usecase.ItemCartResult, err error) {
	filtersCheckProduct := []repository.Filter{
		{
			Column:   "id",
			Value:    param.ProductID,
			Operator: repository.Equality,
		},
		{
			Column:   "deleted_at",
			Operator: repository.IsNULL,
		},
	}
	filtersCheckTcart := &[]repository.Filter{
		{
			Column:   "tic.id",
			Value:    strconv.Itoa(int(id)),
			Operator: repository.Equality,
		},
		{
			Column:   "tic.deleted_at",
			Operator: repository.IsNULL,
		},
	}

	cart, err := c.GetOrCreateCart(ctx, param.UserID)
	if err != nil {
		return
	}

	err = c.cartRepo.StartTx(ctx, pgx.TxOptions{
		IsoLevel:   pgx.ReadCommitted,
		AccessMode: pgx.ReadWrite,
	}, func() error {
		itemCart, err := c.cartRepo.FindOneItemCart(ctx, filtersCheckTcart)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return usecase.ErrProductNotFound
			}
		}

		filtersCheckProduct[0].Value = itemCart.Product.ID
		product, err := c.productRepo.FindOne(ctx, &filtersCheckProduct)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				err = usecase.ErrProductNotFound
			}
			return err
		}

		err = c.cartRepo.UpdateItemCart(ctx, &repository.ItemCart{
			ID:            id,
			CartID:        cart.ID,
			Total:         param.Total,
			SubTotalPrice: param.Total * product.Price,
			Audit: repository.Audit{
				CreatedAt: time.Now().Unix(),
				CreatedBy: param.UserID,
				UpdatedAt: time.Now().Unix(),
			},
		})

		res = &usecase.ItemCartResult{
			ID:            id,
			Total:         param.Total,
			SubTotalPrice: param.Total * product.Price,
			Product: usecase.ProductResult{
				ID:                product.ID,
				Name:              product.Name,
				Price:             product.Price,
				Stock:             product.Stock,
				Description:       product.Description,
				CategoryProductID: product.CategoryProductID,
			},
		}
		return nil
	})

	return
}
