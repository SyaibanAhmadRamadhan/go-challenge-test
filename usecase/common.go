package usecase

type CommonParam struct {
	Device string
	IP     string
	UserID string
}

type PaginateParam struct {
	Page     int
	PageSize int
}

type PaginateResult struct {
	CurrentPage int
	Total       int
	PageSize    int
	PageTotal   int
}
