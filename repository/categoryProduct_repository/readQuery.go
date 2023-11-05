package categoryProduct_repository

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"

	"challenge-test-synapsis/repository"
)

func (c *CategoryProductRepositoryImpl) CheckOne(ctx context.Context, filters *[]repository.Filter) (b bool, err error) {
	filterStr, values, _ := repository.GenerateFilters(filters)
	query := fmt.Sprintf(`SELECT EXISTS (SELECT 1 FROM m_category_product %s)`, filterStr)

	tx, err := c.GetTx()
	if err != nil {
		return
	}

	err = tx.QueryRow(ctx, query, values...).Scan(&b)
	if err != nil {
		return
	}

	return
}

func (c *CategoryProductRepositoryImpl) FindOne(ctx context.Context, filters *[]repository.Filter) (categoryProduct *repository.CategoryProduct, err error) {
	filterStr, values, _ := repository.GenerateFilters(filters)
	query := fmt.Sprintf("SELECT id, name, %s FROM m_category_product %s", repository.AuditToQuery(""), filterStr)

	tx, err := c.GetTx()
	if err != nil {
		return
	}

	categoryProduct = &repository.CategoryProduct{}

	err = tx.QueryRow(ctx, query, values...).Scan(
		&categoryProduct.ID,
		&categoryProduct.Name,
		&categoryProduct.CreatedAt,
		&categoryProduct.CreatedBy,
		&categoryProduct.UpdatedAt,
		&categoryProduct.UpdatedBy,
		&categoryProduct.DeletedAt,
		&categoryProduct.DeletedBy,
	)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			log.Warn().Msgf("failed query row | err: %v", err)
		}
		categoryProduct = nil
		return
	}

	return
}

func (c *CategoryProductRepositoryImpl) FindAllAndSearch(
	ctx context.Context, param repository.FindAllAndSearchParam,
) (categoryProducts *[]repository.CategoryProduct, total int, err error) {
	filterStr, values, lastPH := repository.GenerateFilters(param.Filters)
	orderStr := param.Pagination.GenerateOrderBy()

	tx, err := c.GetTx()
	if err != nil {
		return
	}

	search := ""
	if param.Filters != nil {
		search += "AND "
	}
	search += "name LIKE $" + strconv.Itoa(lastPH)
	lastPH++

	values = append(values, "%"+param.Search+"%")

	queryCount := fmt.Sprintf("SELECT COUNT(*) FROM m_category_product %s %s", filterStr, search)

	err = tx.QueryRow(ctx, queryCount, values...).Scan(&total)
	if err != nil {
		return
	}

	query := fmt.Sprintf("SELECT id, name, %s FROM m_category_product %s %s %s LIMIT $%d OFFSET $%d",
		repository.AuditToQuery(""), filterStr, search, orderStr, lastPH, lastPH+1)

	values = append(values, param.Pagination.Limit)
	values = append(values, param.Pagination.Offset)

	rows, err := tx.Query(ctx, query, values...)
	if err != nil {
		return
	}

	categoryProduct := &repository.CategoryProduct{}
	categoryProducts = &[]repository.CategoryProduct{}

	for rows.Next() {
		err = rows.Scan(
			&categoryProduct.ID,
			&categoryProduct.Name,
			&categoryProduct.CreatedAt,
			&categoryProduct.CreatedBy,
			&categoryProduct.UpdatedAt,
			&categoryProduct.UpdatedBy,
			&categoryProduct.DeletedAt,
			&categoryProduct.DeletedBy,
		)
		if err != nil {
			return
		}

		*categoryProducts = append(*categoryProducts, *categoryProduct)
	}

	return
}
