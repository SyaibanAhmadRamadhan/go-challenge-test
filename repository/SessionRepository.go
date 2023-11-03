package repository

type Session struct {
	ID      int    `sql:"id"`
	Token   string `sql:"token"`
	Device  string `sql:"device"`
	LoginAt int    `sql:"login_at"`
	IP      int    `sql:"ip"`
	Audit
	User *User
}

type SessionRepository interface {
	UOWRepository
}
