package repository

import (
	"belajar-rest-api-golang/model/domain"
	"context"
	"database/sql"
)

/*
membuat kontrak dalam bentuk interface terlebih dahulu
susunannya yaitu nama function(context, sql transaction/tidak, model) return value nya
*/

type ProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	Delete(ctx context.Context, tx *sql.Tx, product domain.Product)
	FindById(ctx context.Context, tx *sql.Tx, productId int) (domain.Product, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Product
}
