package app

import (
	"belajar-rest-api-golang/controller"
	"belajar-rest-api-golang/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(productController controller.ProductController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/products", productController.FindAll)
	router.GET("/api/products/:productId", productController.FindById)
	router.POST("/api/products", productController.Create)
	router.PUT("/api/products/:productId", productController.Update)
	router.DELETE("/api/products/:productId", productController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
