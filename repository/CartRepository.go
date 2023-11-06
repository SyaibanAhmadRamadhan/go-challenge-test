package repository

import (
	"context"
)

type Cart struct {
	ID     int64  `sql:"id"`
	UserID string `sql:"user_id"`
	Audit
	ItemCarts *[]ItemCart
}

// ItemCart is Transaction Cart
type ItemCart struct {
	ID            int64 `sql:"id"`
	CartID        int64 `sql:"cart_id"`
	Total         int   `sql:"total"`
	SubTotalPrice int   `sql:"sub_total_price"`
	Audit
	Product *Product
}

type CartRepository interface {
	CreateItemCart(ctx context.Context, itemCart *ItemCart) (id int64, err error)
	UpdateItemCart(ctx context.Context, itemCart *ItemCart) (err error)
	DeleteItemCart(ctx context.Context, id int64, userID string) (err error)
	CheckOneItemCart(ctx context.Context, filters *[]Filter) (b bool, err error)
	FindOneItemCart(ctx context.Context, filters *[]Filter) (itemCart *ItemCart, err error)
	FindAllItemCartByCart(ctx context.Context, param *FPSParam) (cart *Cart, total int, err error)
	Create(ctx context.Context, cart *Cart) (id int64, err error)
	Delete(ctx context.Context, id string, userID string) (err error)
	CheckOne(ctx context.Context, filters *[]Filter) (b bool, err error)
	FindOne(ctx context.Context, filters *[]Filter) (cart *Cart, err error)
	UOWRepository
}
