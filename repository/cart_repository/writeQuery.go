package cart_repository

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"challenge-test-synapsis/repository"
)

func (c *CartRepositoryImpl) CreateItemCart(ctx context.Context, cart *repository.ItemCart) (id int64, err error) {
	query := `INSERT INTO t_item_cart (cart_id, product_id, total, sub_total_price, created_at, created_by, updated_at) 
					VALUES ($1, $2, $3, $4, $5, $6, $7)`

	tx, err := c.GetTx()
	if err != nil {
		return
	}

	res, err := tx.Exec(ctx, query,
		cart.CartID,
		cart.Product.ID,
		cart.Total,
		cart.SubTotalPrice,
		cart.CreatedAt,
		cart.CreatedBy,
		cart.UpdatedAt,
	)
	if err != nil {
		log.Warn().Msgf("failed exec command | err: %v", err)
		return
	}

	id = res.RowsAffected()
	return
}

func (c *CartRepositoryImpl) UpdateItemCart(ctx context.Context, itemCart *repository.ItemCart) (err error) {
	query := `UPDATE t_item_cart SET total=$1, sub_total_price=$2, updated_at=$3, updated_by=$4 WHERE id=$5 AND cart_id=$6 AND deleted_at IS NULL`

	tx, err := c.GetTx()
	if err != nil {
		return
	}

	res, err := tx.Exec(ctx, query,
		itemCart.Total,
		itemCart.SubTotalPrice,
		itemCart.UpdatedAt,
		itemCart.UpdatedBy,
		itemCart.ID,
		itemCart.CartID,
	)
	if err != nil {
		log.Warn().Msgf("failed exec command | err: %v", err)
		return
	}

	if res.RowsAffected() == 0 {
		log.Info().Msgf("updated does not meet the conditions")
	}

	return
}

func (c *CartRepositoryImpl) DeleteItemCart(ctx context.Context, id int64, userID string) (err error) {
	query := `UPDATE t_item_cart SET deleted_at=$1, deleted_by=$2 WHERE id=$3 AND deleted_at IS NULL`

	tx, err := c.GetTx()
	if err != nil {
		return
	}

	res, err := tx.Exec(ctx, query,
		time.Now().Unix(),
		userID,
		id,
	)
	if err != nil {
		log.Warn().Msgf("failed exec command | err: %v", err)
		return
	}

	if res.RowsAffected() == 0 {
		log.Info().Msgf("deleted does not meet the conditions")
	}

	return
}

func (c *CartRepositoryImpl) Create(ctx context.Context, cart *repository.Cart) (id int64, err error) {
	query := `INSERT INTO m_cart (user_id, created_at, created_by, updated_at) 
					VALUES ($1, $2, $3, $4, $5)`

	tx, err := c.GetTx()
	if err != nil {
		return
	}

	res, err := tx.Exec(ctx, query,
		cart.UserID,
		cart.CreatedAt,
		cart.CreatedBy,
		cart.UpdatedAt,
	)
	if err != nil {
		log.Warn().Msgf("failed exec command | err: %v", err)
		return
	}

	id = res.RowsAffected()
	return
}

func (c *CartRepositoryImpl) Delete(ctx context.Context, id string, userID string) (err error) {
	query := `UPDATE m_cart SET deleted_at=$1, deleted_by=$2 WHERE id=$3 AND deleted_at IS NULL`

	tx, err := c.GetTx()
	if err != nil {
		return
	}

	res, err := tx.Exec(ctx, query,
		time.Now().Unix(),
		userID,
		id,
	)
	if err != nil {
		log.Warn().Msgf("failed exec command | err: %v", err)
		return
	}

	if res.RowsAffected() == 0 {
		log.Info().Msgf("updated does not meet the conditions")
	}

	return
}
