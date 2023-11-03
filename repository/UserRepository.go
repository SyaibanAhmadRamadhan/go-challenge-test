package repository

import (
	"context"
)

type User struct {
	Audit
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id string) error
	CheckOne(ctx context.Context, filters *[]Filter) (b bool, err error)
	FindOne(ctx context.Context, filters *[]Filter) (user *User, err error)
	FindAll(ctx context.Context, filters *[]Filter, paginate Pagination) (users *[]User, total int, err error)
	Search(ctx context.Context, search SearchParam) (users *[]User, total int, err error)
}
