package modules

import (
	"net/http"

	"github.com/test-tecnico-grupo-lagos/backend/controllers"
)

func HandleUsers(w http.ResponseWriter, r *http.Request) error {
	handler := controllers.NewUserController()
	if r.Method == http.MethodPost {
		return handler.HandleCreateUser(w, r)
	}
	return http.ErrNotSupported
}

func HandleUsersWithParams(w http.ResponseWriter, r *http.Request) error {
	handler := controllers.NewUserController()
	if r.Method == http.MethodGet {
		return handler.HandleFindUserByID(w, r)
	}
	return http.ErrNotSupported
}
