package modules

import (
	"net/http"

	"github.com/test-tecnico-grupo-lagos/backend/controllers"
)

func HandleFindCart(w http.ResponseWriter, r *http.Request) error {
	handler := controllers.NewCartController()
	if r.Method == http.MethodGet {
		return handler.HandleUpdateCart(w, r)
	}
	return http.ErrNotSupported
}
