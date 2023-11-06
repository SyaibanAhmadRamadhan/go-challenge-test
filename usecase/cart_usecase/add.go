package cart_usecase

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"

	"challenge-test-synapsis/repository"
	"challenge-test-synapsis/usecase"
)

func (c *CartUsecaseImpl) AddItemCart(ctx context.Context, param *usecase.ItemCartParam) (res *usecase.ItemCartResult, err error) {
	filtersCheckProduct := &[]repository.Filter{
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
	filtersCheckItemCart := &[]repository.Filter{
		{
			Column:   "tic.product_id",
			Value:    param.ProductID,
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
	var id int64

	err = c.cartRepo.StartTx(ctx, pgx.TxOptions{
		IsoLevel:   pgx.ReadCommitted,
		AccessMode: pgx.ReadWrite,
	}, func() error {
		itemCart, err := c.cartRepo.FindOneItemCart(ctx, filtersCheckItemCart)
		if err != nil {
			if !errors.Is(err, pgx.ErrNoRows) {
				return err
			}
		}

		product, err := c.productRepo.FindOne(ctx, filtersCheckProduct)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				err = usecase.ErrProductNotFound
			}
			return err
		}

		if itemCart != nil {
			err = c.cartRepo.UpdateItemCart(ctx, &repository.ItemCart{
				ID:            itemCart.ID,
				CartID:        cart.ID,
				Total:         param.Total,
				SubTotalPrice: param.Total * product.Price,
				Audit: repository.Audit{
					CreatedAt: time.Now().Unix(),
					CreatedBy: param.UserID,
					UpdatedAt: time.Now().Unix(),
				},
				Product: product,
			})
			id = itemCart.ID
		} else {
			idRes, err := c.cartRepo.CreateItemCart(ctx, &repository.ItemCart{
				CartID:        cart.ID,
				Total:         param.Total,
				SubTotalPrice: param.Total * product.Price,
				Audit: repository.Audit{
					CreatedAt: time.Now().Unix(),
					CreatedBy: param.UserID,
					UpdatedAt: time.Now().Unix(),
				},
				Product: product,
			})
			if err != nil {
				return err
			}
			id = idRes
		}

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
		return err
	})

	return
}
