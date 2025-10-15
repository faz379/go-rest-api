package web

/*
product response untuk response yang dikirim ke client
sehingga jika ada data yang sensitif tidak perlu dikirim ke client
*/
type ProductResponse struct {
	Id          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	Description string  `json:"description"`
}
