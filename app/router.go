package app

import (
	"belajar-rest-api-golang/controller"
	"belajar-rest-api-golang/exception"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(productController controller.ProductController, userController *controller.UserControllerImpl) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/users/register", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		userController.Register(writer, request)
	})
	router.POST("/api/users/login", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		userController.Login(writer, request)
	})
	router.POST("/api/users/logout", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		userController.Logout(writer, request)
	})

	router.GET("/api/products", productController.FindAll)
	router.GET("/api/products/:productId", productController.FindById)
	router.POST("/api/products", productController.Create)
	router.PUT("/api/products/:productId", productController.Update)
	router.DELETE("/api/products/:productId", productController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
