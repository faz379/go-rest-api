package service

import (
	"belajar-rest-api-golang/model/web"
	"context"
)

/*
membuat kontrak dalam bentuk interface terlebih dahulu
di service ini berupa logic bisnis dari aplikasinya
*/
type ProductService interface {
	Create(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse
	Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductResponse
	Delete(ctx context.Context, productId int)
	FindById(ctx context.Context, productId int) web.ProductResponse
	FindAll(ctx context.Context) []web.ProductResponse
}
