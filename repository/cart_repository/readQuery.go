package cart_repository

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"

	"challenge-test-synapsis/repository"
)

func (c *CartRepositoryImpl) CheckOneItemCart(ctx context.Context, filters *[]repository.Filter) (b bool, err error) {
	filterStr, values, _ := repository.GenerateFilters(filters)
	query := fmt.Sprintf(`SELECT EXISTS (SELECT 1 FROM t_item_cart %s)`, filterStr)

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

func (c *CartRepositoryImpl) FindOneItemCart(ctx context.Context, filters *[]repository.Filter) (itemCart *repository.ItemCart, err error) {
	filterStr, values, _ := repository.GenerateFilters(filters)
	auditQueryTC := repository.AuditToQuery("tic.")
	auditQueryMP := repository.AuditToQuery("mp.")
	query := fmt.Sprintf(`SELECT tic.id, tic.cart_id, tic.total, tic.sub_total_price, %s, 
										 mp.id, mp.category_product_id, mp.name, mp.stock, mp.price, mp.description, %s
       							  FROM t_item_cart as tic 
								  LEFT JOIN m_product as mp ON tic.product_id = mp.id
       							  %s`,
		auditQueryTC, auditQueryMP, filterStr)

	tx, err := c.GetTx()
	if err != nil {
		return
	}

	itemCart = &repository.ItemCart{
		Product: &repository.Product{},
	}

	err = tx.QueryRow(ctx, query, values...).Scan(
		&itemCart.ID, &itemCart.CartID, &itemCart.Total, &itemCart.SubTotalPrice, &itemCart.CreatedAt, &itemCart.CreatedBy, &itemCart.UpdatedAt,
		&itemCart.UpdatedBy, &itemCart.DeletedAt, &itemCart.DeletedBy,
		&itemCart.Product.ID, &itemCart.Product.CategoryProductID, &itemCart.Product.Name, &itemCart.Product.Stock,
		&itemCart.Product.Price, &itemCart.Product.Description, &itemCart.Product.CreatedAt, &itemCart.Product.CreatedBy,
		&itemCart.Product.UpdatedAt, &itemCart.Product.UpdatedBy, &itemCart.Product.DeletedAt, &itemCart.Product.DeletedBy,
	)

	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			log.Warn().Msgf("failed query row | err: %v", err)
		}
		itemCart = nil
		return
	}

	return
}

func (c *CartRepositoryImpl) FindAllItemCartByCart(ctx context.Context, param *repository.FPSParam) (cart *repository.Cart, total int, err error) {
	filterStr, values, lastPH := repository.GenerateFilters(param.Filters)
	auditQueryMC := repository.AuditToQuery("mc.")
	auditQueryTC := repository.AuditToQuery("tic.")
	auditQueryMP := repository.AuditToQuery("mp.")
	orderStr := param.Pagination.GenerateOrderBy()

	tx, err := c.GetTx()
	if err != nil {
		return
	}

	search := ""
	if param.Filters != nil {
		search += "AND "
	}
	search += "mp.name LIKE $" + strconv.Itoa(lastPH)
	lastPH++

	values = append(values, "%"+param.Search+"%")

	queryCount := fmt.Sprintf(`SELECT COUNT(*) FROM m_cart AS mc 
    										LEFT JOIN t_item_cart AS tic ON mc.id = tic.cart_id 
											LEFT JOIN m_product AS mp ON tic.product_id = mp.id 
    										%s %s`, filterStr, search)

	err = tx.QueryRow(ctx, queryCount, values...).Scan(&total)
	if err != nil {
		return
	}

	query := fmt.Sprintf(`SELECT mc.id, mc.user_id, %s,
    									tic.id, tic.total, tic.sub_total_price, %s, 
										mp.id, mp.category_product_id, mp.name, mp.stock, mp.price, mp.description, %s
       							  FROM m_cart as mc 
								  LEFT JOIN t_item_cart as tic ON mc.id = tic.cart_id
								  LEFT JOIN m_product as mp ON tic.product_id = mp.id
       							  %s %s %s LIMIT $%d OFFSET $%d`,
		auditQueryMC, auditQueryTC, auditQueryMP, filterStr, search, orderStr, lastPH, lastPH+1)

	values = append(values, param.Pagination.Limit)
	values = append(values, param.Pagination.Offset)

	rows, err := tx.Query(ctx, query, values...)
	if err != nil {
		log.Warn().Msgf("failed query data | err:%v", err)
		return
	}

	cart = &repository.Cart{
		ItemCarts: &[]repository.ItemCart{},
	}

	for rows.Next() {
		itemCart := repository.ItemCart{
			Product: &repository.Product{},
		}
		err = rows.Scan(
			&cart.ID, &cart.UserID, &cart.CreatedAt, &cart.CreatedBy, &cart.UpdatedAt, &cart.UpdatedBy,
			&cart.DeletedAt, &cart.DeletedBy,
			&itemCart.ID, &itemCart.Total, &itemCart.SubTotalPrice, &itemCart.CreatedAt, &itemCart.CreatedBy,
			&itemCart.UpdatedAt, &itemCart.UpdatedBy, &itemCart.DeletedAt, &itemCart.DeletedBy,
			&itemCart.Product.ID, &itemCart.Product.CategoryProductID, &itemCart.Product.Name, &itemCart.Product.Stock,
			&itemCart.Product.Price, &itemCart.Product.Description, &itemCart.Product.CreatedAt, &itemCart.Product.CreatedBy,
			&itemCart.Product.UpdatedAt, &itemCart.Product.UpdatedBy, &itemCart.Product.DeletedAt,
			&itemCart.Product.DeletedBy,
		)
		if err != nil {
			log.Warn().Msgf("failed scan query | err: %v", err)
			cart = nil
			return
		}

		*cart.ItemCarts = append(*cart.ItemCarts, itemCart)
	}

	return
}

func (c *CartRepositoryImpl) CheckOne(ctx context.Context, filters *[]repository.Filter) (b bool, err error) {
	filterStr, values, _ := repository.GenerateFilters(filters)
	query := fmt.Sprintf(`SELECT EXISTS (SELECT 1 FROM m_cart %s)`, filterStr)

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

func (c *CartRepositoryImpl) FindOne(ctx context.Context, filters *[]repository.Filter) (cart *repository.Cart, err error) {
	filterStr, values, _ := repository.GenerateFilters(filters)
	auditQuery := repository.AuditToQuery("")

	tx, err := c.GetTx()
	if err != nil {
		return
	}

	query := fmt.Sprintf(`SELECT id, user_id, %s
       							  FROM m_cart %s`,
		auditQuery, filterStr)

	cart = &repository.Cart{}

	err = tx.QueryRow(ctx, query, values...).Scan(
		&cart.ID, &cart.UserID, &cart.CreatedAt, &cart.CreatedBy, &cart.UpdatedAt, &cart.UpdatedBy,
		&cart.DeletedAt, &cart.DeletedBy,
	)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			log.Warn().Msgf("failed query row | err: %v", err)
		}
		cart = nil
	}

	return
}
