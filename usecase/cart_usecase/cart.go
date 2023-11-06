package cart_usecase

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"

	"challenge-test-synapsis/repository"
	"challenge-test-synapsis/usecase"
)

type CartUsecaseImpl struct {
	cartRepo    repository.CartRepository
	productRepo repository.ProductRepository
}

func NewCartUsecaseImpl(
	cartRepo repository.CartRepository,
	productRepo repository.ProductRepository,
) usecase.CartUsecase {
	return &CartUsecaseImpl{
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

func (c *CartUsecaseImpl) GetOrCreateCart(ctx context.Context, userID string) (cart *repository.Cart, err error) {
	filter := &[]repository.Filter{
		{
			Column:   "user_id",
			Value:    userID,
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
		cart, err = c.cartRepo.FindOne(ctx, filter)
		if err != nil {
			if !errors.Is(err, pgx.ErrNoRows) {
				return err
			}
		}

		if cart == nil {
			cart = &repository.Cart{
				UserID: userID,
				Audit: repository.Audit{
					CreatedAt: time.Now().Unix(),
					CreatedBy: userID,
					UpdatedAt: time.Now().Unix(),
				},
				ItemCarts: nil,
			}
			id, err := c.cartRepo.Create(ctx, cart)
			if err != nil {
				return err
			}
			cart.ID = id
		}

		return nil
	})

	return
}
