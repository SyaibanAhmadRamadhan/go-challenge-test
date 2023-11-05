package product_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"

	"challenge-test-synapsis/repository"
)

func (p *ProductRepositoryImpl) CheckOne(ctx context.Context, filters *[]repository.Filter) (b bool, err error) {
	filterStr, values, _ := repository.GenerateFilters(filters)
	query := fmt.Sprintf(`SELECT EXISTS (SELECT 1 FROM m_product %s)`, filterStr)

	tx, err := p.GetTx()
	if err != nil {
		return
	}

	err = tx.QueryRow(ctx, query, values...).Scan(&b)
	if err != nil {
		return
	}

	return
}

func (p *ProductRepositoryImpl) FindOne(ctx context.Context, filters *[]repository.Filter) (product *repository.Product, err error) {
	filterStr, values, _ := repository.GenerateFilters(filters)
	query := fmt.Sprintf("SELECT id, name, stock, price, description, %s FROM m_product %s", repository.AuditToQuery(""), filterStr)

	tx, err := p.GetTx()
	if err != nil {
		return
	}

	product = &repository.Product{}

	err = tx.QueryRow(ctx, query, values...).Scan(
		&product.ID,
		&product.Name,
		&product.Stock,
		&product.Price,
		&product.Description,
		&product.CreatedAt,
		&product.CreatedBy,
		&product.UpdatedAt,
		&product.UpdatedBy,
		&product.DeletedAt,
		&product.DeletedBy,
	)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			log.Warn().Msgf("failed query row | err: %v", err)
		}
		product = nil
		return
	}

	return
}

func (p *ProductRepositoryImpl) FindAllAndSearch(
	ctx context.Context, param repository.FindAllAndSearchParam,
) (products *[]repository.Product, total int, err error) {
	filterStr, values, lastPH := repository.GenerateFilters(param.Filters)
	orderStr := param.Pagination.GenerateOrderBy()

	tx, err := p.GetTx()
	if err != nil {
		return
	}

	search := ""
	if param.Filters != nil {
		search += "AND "
	}
	search += fmt.Sprintf("(name LIKE $%d OR description LIKE $%d)", lastPH, lastPH)
	lastPH++

	values = append(values, "%"+param.Search+"%")

	queryCount := fmt.Sprintf("SELECT COUNT(*) FROM m_product %s %s", filterStr, search)

	err = tx.QueryRow(ctx, queryCount, values...).Scan(&total)
	if err != nil {
		return
	}

	query := fmt.Sprintf("SELECT id, name, stock, price, description, %s FROM m_product %s %s %s LIMIT $%d OFFSET $%d",
		repository.AuditToQuery(""), filterStr, search, orderStr, lastPH, lastPH+1)

	values = append(values, param.Pagination.Limit)
	values = append(values, param.Pagination.Offset)

	rows, err := tx.Query(ctx, query, values...)
	if err != nil {
		return
	}

	product := &repository.Product{}
	products = &[]repository.Product{}

	for rows.Next() {
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Stock,
			&product.Price,
			&product.Description,
			&product.CreatedAt,
			&product.CreatedBy,
			&product.UpdatedAt,
			&product.UpdatedBy,
			&product.DeletedAt,
			&product.DeletedBy,
		)
		if err != nil {
			return
		}

		*products = append(*products, *product)
	}

	return
}
