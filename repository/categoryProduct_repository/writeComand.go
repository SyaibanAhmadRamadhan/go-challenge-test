package categoryProduct_repository

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"challenge-test-synapsis/repository"
)

func (c *CategoryProductRepositoryImpl) Create(ctx context.Context, categoryProduct *repository.CategoryProduct) (err error) {
	query := `INSERT INTO m_category_product (id, name, created_at, created_by, updated_at) VALUES ($1, $2, $3, $4, $5)`

	tx, err := c.GetTx()
	if err != nil {
		return
	}

	_, err = tx.Exec(ctx, query,
		categoryProduct.ID,
		categoryProduct.Name,
		categoryProduct.CreatedAt,
		categoryProduct.CreatedBy,
		categoryProduct.UpdatedAt,
	)
	if err != nil {
		log.Warn().Msgf("failed exec command | err: %v", err)
		return
	}

	return
}

func (c *CategoryProductRepositoryImpl) Update(ctx context.Context, categoryProduct *repository.CategoryProduct) (err error) {
	query := `UPDATE m_category_product SET name=$1, updated_at=$2, updated_by=$3 WHERE id=$4 AND deleted_at IS NULL`

	tx, err := c.GetTx()
	if err != nil {
		return
	}

	res, err := tx.Exec(ctx, query,
		categoryProduct.Name,
		categoryProduct.UpdatedAt,
		categoryProduct.UpdatedBy,
		categoryProduct.ID,
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

func (c *CategoryProductRepositoryImpl) Delete(ctx context.Context, id string, userID string) (err error) {
	query := `UPDATE m_category_product SET deleted_at=$1, deleted_by=$2 WHERE id=$3 AND deleted_at IS NULL`

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
