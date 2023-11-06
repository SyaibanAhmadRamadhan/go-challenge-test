package usecase

import (
	"context"

	"challenge-test-synapsis/repository"
)

type ItemCartParam struct {
	ProductID string
	Total     int
	CommonParam
}

type GetCartParam struct {
	Search string
	PaginateParam
	CommonParam
}

type CartResult struct {
	ID        int64             `json:"id"`
	ItemCarts *[]ItemCartResult `json:"item_carts"`
}

const StatusProductNotAvailable = "product is not available"
const StatusProductAvailable = "product is available"
const StatusProductEmpty = "product stock is empty"
const StatusTotalExceedStock = "Total product Exceeds Stock"

type ItemCartResult struct {
	ID            int64         `json:"id"`
	Total         int           `json:"total"`
	SubTotalPrice int           `json:"sub_total_price"`
	StatusProduct string        `json:"status_product"`
	Product       ProductResult `json:"product"`
}

type CartUsecase interface {
	AddItemCart(ctx context.Context, param *ItemCartParam) (res *ItemCartResult, err error)
	UpdateItemCart(ctx context.Context, id int64, param *ItemCartParam) (res *ItemCartResult, err error)
	DeleteItemCart(ctx context.Context, id int64, param *CommonParam) (err error)
	GetItemCartByCart(ctx context.Context, param *GetCartParam) (res *CartResult, paginate *PaginateResult, err error)
	GetOrCreateCart(ctx context.Context, userID string) (cart *repository.Cart, err error)
}
