package middleware

import (
	"belajar-rest-api-golang/helper"
	"belajar-rest-api-golang/model/web"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path

	// Path yang tidak butuh Authorization
	publicPaths := []string{
		"/api/users/login",
		"/api/users/register",
	}

	for _, p := range publicPaths {
		if strings.HasPrefix(path, p) {
			middleware.Handler.ServeHTTP(writer, request)
			return
		}
	}

	// Ambil header Authorization
	authHeader := request.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	// Jika valid, lanjutkan ke handler berikutnya
	middleware.Handler.ServeHTTP(writer, request)
}
