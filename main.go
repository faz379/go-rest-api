package main

import (
	"belajar-rest-api-golang/app"
	"belajar-rest-api-golang/controller"
	"belajar-rest-api-golang/helper"
	"belajar-rest-api-golang/middleware"
	"belajar-rest-api-golang/repository"
	"belajar-rest-api-golang/service"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db, validate)
	productController := controller.NewProductController(productService)

	router := app.NewRouter(productController)

	server := http.Server{
		Addr:    "localhost:5000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	fmt.Println("Server running on:", server.Addr)
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
