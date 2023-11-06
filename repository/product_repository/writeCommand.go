package product_repository

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"challenge-test-synapsis/repository"
)

func (p *ProductRepositoryImpl) Create(ctx context.Context, product *repository.Product) (err error) {
	query := `INSERT INTO m_product (id, category_product_id, name, stock, price, description, created_at, created_by, updated_at) 
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	tx, err := p.GetTx()
	if err != nil {
		return
	}

	_, err = tx.Exec(ctx, query,
		product.ID,
		product.CategoryProductID,
		product.Name,
		product.Stock,
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
	query := `UPDATE m_product SET category_product_id=$1, name=$2, stock=$3, price=$4, description=$5, updated_at=$6, updated_by=$7 
                 WHERE id=$8 AND deleted_at IS NULL`

	tx, err := p.GetTx()
	if err != nil {
		return
	}

	res, err := tx.Exec(ctx, query,
		product.CategoryProductID,
		product.Name,
		product.Stock,
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
