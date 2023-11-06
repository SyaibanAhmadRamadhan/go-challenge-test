package cart_usecase

import (
	"context"
	"math"
	"strconv"

	"github.com/jackc/pgx/v5"

	"challenge-test-synapsis/repository"
	"challenge-test-synapsis/usecase"
)

func (c *CartUsecaseImpl) GetItemCartByCart(ctx context.Context, param *usecase.GetCartParam) (res *usecase.CartResult, paginate *usecase.PaginateResult, err error) {
	cart, err := c.GetOrCreateCart(ctx, param.UserID)
	if err != nil {
		return
	}
	err = c.cartRepo.StartTx(ctx, pgx.TxOptions{
		IsoLevel:   pgx.ReadCommitted,
		AccessMode: pgx.ReadWrite,
	}, func() error {
		filters := []repository.Filter{
			{
				Prefix:   "tic.",
				Column:   "deleted_at",
				Operator: repository.IsNULL,
			},
		}
		filters = append(filters, repository.Filter{
			Prefix:   "mc.",
			Column:   "id",
			Value:    strconv.Itoa(int(cart.ID)),
			Operator: repository.Equality,
		})

		findAllAndSearch := repository.FPSParam{
			Filters: &filters,
			Pagination: repository.Pagination{
				Limit:  param.PageSize,
				Offset: (param.Page - 1) * param.PageSize,
				Orders: map[string]string{
					"tic.id": "DESC",
				},
			},
			Search: param.Search,
		}

		data, total, err := c.cartRepo.FindAllItemCartByCart(ctx, &findAllAndSearch)

		paginate = &usecase.PaginateResult{}
		paginate.CurrentPage = param.Page
		paginate.Total = total
		paginate.PageTotal = int(math.Ceil(float64(paginate.Total) / float64(param.PageSize)))
		paginate.PageSize = param.PageSize

		cart = data
		return err
	})

	if err != nil {
		return
	}
	res = &usecase.CartResult{
		ID:        cart.ID,
		ItemCarts: &[]usecase.ItemCartResult{},
	}

	status := usecase.StatusProductAvailable
	for _, itemCart := range *cart.ItemCarts {
		if itemCart.Product.DeletedAt.Valid {
			status = usecase.StatusProductNotAvailable
		} else if itemCart.Product.Stock <= 0 {
			status = usecase.StatusProductEmpty
		} else if itemCart.Product.Stock < itemCart.Total {
			status = usecase.StatusTotalExceedStock
		}
		*res.ItemCarts = append(*res.ItemCarts, usecase.ItemCartResult{
			ID:            itemCart.ID,
			Total:         itemCart.Total,
			SubTotalPrice: itemCart.Total * itemCart.Product.Price,
			StatusProduct: status,
			Product: usecase.ProductResult{
				ID:                itemCart.Product.ID,
				Name:              itemCart.Product.Name,
				Price:             itemCart.Product.Price,
				Stock:             itemCart.Product.Stock,
				Description:       itemCart.Product.Description,
				CategoryProductID: itemCart.Product.CategoryProductID,
			},
		})
	}

	return
}
