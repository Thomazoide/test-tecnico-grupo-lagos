package modules

import (
	"net/http"

	"github.com/test-tecnico-grupo-lagos/backend/controllers"
)

func HandleGetProductsByPage(w http.ResponseWriter, r *http.Request) error {
	handler := controllers.NewProductsController()
	if r.Method == http.MethodGet {
		return handler.HandleGetProductsFromPage(w, r)
	}
	return http.ErrNotSupported
}

func HandleGetProductByCode(w http.ResponseWriter, r *http.Request) error {
	handler := controllers.NewProductsController()
	if r.Method == http.MethodGet {
		return handler.HandleGetProductFromCode(w, r)
	}
	return http.ErrNotSupported
}
