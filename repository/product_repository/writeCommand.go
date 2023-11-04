package product_repository

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"challenge-test-synapsis/repository"
)

func (p *ProductRepositoryImpl) Create(ctx context.Context, product *repository.Product) (err error) {
	query := `INSERT INTO m_product (id, name, price, description, created_at, created_by, updated_at) 
					VALUES ($1, $2, $3, $4, $5, $6, $7)`

	tx, err := p.GetTx()
	if err != nil {
		return
	}

	_, err = tx.Exec(ctx, query,
		product.ID,
		product.Name,
		product.Price,
		product.Description,
		product.CreatedAt,
		product.CreatedBy,
		product.UpdatedAt,
	)
	if err != nil {
		log.Warn().Msgf("failed exec command | err: %v", err)
		return
	}

	return
}

func (p *ProductRepositoryImpl) Update(ctx context.Context, product *repository.Product) (err error) {
	query := `UPDATE m_product SET name=$1, price=$2, description=$3, updated_at=$4, updated_by=$5 WHERE id=$6 AND deleted_at IS NULL`

	tx, err := p.GetTx()
	if err != nil {
		return
	}

	res, err := tx.Exec(ctx, query,
		product.Name,
		product.Price,
		product.Description,
		product.UpdatedAt,
		product.UpdatedBy,
		product.ID,
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

func (p *ProductRepositoryImpl) Delete(ctx context.Context, id string, userID string) (err error) {
	query := `UPDATE m_product SET deleted_at=$1, deleted_by=$2 WHERE id=$3 AND deleted_at IS NULL`

	tx, err := p.GetTx()
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
