package repository

import (
	"context"
)

type Session struct {
	ID      string `sql:"id"`
	UserID  string `sql:"user_id"`
	Token   string `sql:"token"`
	Device  string `sql:"device"`
	LoginAt int64  `sql:"login_at"`
	IP      string `sql:"ip"`
	Audit
}

type SessionRepository interface {
	Create(ctx context.Context, session *Session) (err error)
	Update(ctx context.Context, session *Session) (err error)
	Delete(ctx context.Context, id string, userID string) (err error)
	CheckOne(ctx context.Context, filters *[]Filter) (b bool, err error)
	FindOne(ctx context.Context, filters *[]Filter) (session *Session, err error)
	UOWRepository
}
