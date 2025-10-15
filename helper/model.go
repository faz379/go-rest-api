package helper

import (
	"belajar-rest-api-golang/model/domain"
	"belajar-rest-api-golang/model/web"
)

func ToProductResponse(product domain.Product) web.ProductResponse {
	return web.ProductResponse{
		Id:          product.Id,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
	}
}

func ToProductResponses(products []domain.Product) []web.ProductResponse {
	var productResponses []web.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}
	return productResponses
}
