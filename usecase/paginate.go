package usecase

type PaginateParam struct {
	Page      int
	Total     int
	PageSize  int
	PageTotal [5]int
}
