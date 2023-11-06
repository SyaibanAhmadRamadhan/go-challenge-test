package schema

type ProductResponse struct {
	ID                string `json:"id"`
	CategoryProductID string `json:"category_product_id"`
	Name              string `json:"name"`
	Price             int    `json:"price"`
	Stock             int    `json:"stock"`
	Description       string `json:"description"`
}
