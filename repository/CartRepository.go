package repository

type Cart struct {
	ID         string `sql:"id"`
	TotalPrice int    `sql:"total_price"`
	Audit
	User  *User
	TCart *[]TCart
}

type TCart struct {
	ID            int    `sql:"id"`
	CartID        string `sql:"cart_id"`
	Total         int    `sql:"total"`
	SubTotalPrice int    `sql:"sub_total_price"`
	Audit
	Product *Product
}

type CartRepository interface {
	UOWRepository
}
