package web

type ProductCreateRequest struct {
	Name        string  `json:"name" validate:"required"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	Description string  `json:"description"`
}
