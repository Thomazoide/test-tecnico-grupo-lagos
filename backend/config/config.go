package config

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/test-tecnico-grupo-lagos/backend/middleware"
	"github.com/test-tecnico-grupo-lagos/backend/modules"
	"github.com/test-tecnico-grupo-lagos/backend/utils"
)

type APIServer struct {
	listenAddress string
}

func NewApiServer(listenAddress string) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
	}
}

// constantes que seran para definir las rutas del server
const (
	UsersRoute          string = "/users"
	UserByIDRoute       string = "/users/{id}"
	UpdateCartRoute     string = "/cart/{id}"
	ProductsByPageRoute string = "/products/page/{page}"
	ProductsByCodeRoute string = "/products/code/{code}"
)

func (server *APIServer) RunServer() {
	router := mux.NewRouter()
	middleware.EnableCORS(router)
	router.Use(middleware.MiddlewareCORS)
	router.HandleFunc(UsersRoute, utils.MakeHTTPHandlerFunc(modules.HandleUsers))
	router.HandleFunc(UserByIDRoute, utils.MakeHTTPHandlerFunc(modules.HandleUsersWithParams))
	router.HandleFunc(UpdateCartRoute, utils.MakeHTTPHandlerFunc(modules.HandleUpdateCart))
	router.HandleFunc(ProductsByPageRoute, utils.MakeHTTPHandlerFunc(modules.HandleGetProductsByPage))
	router.HandleFunc(ProductsByCodeRoute, utils.MakeHTTPHandlerFunc(modules.HandleGetProductByCode))
	http.ListenAndServe(server.listenAddress, router)
}
