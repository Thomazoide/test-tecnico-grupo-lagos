package config

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/test-tecnico-grupo-lagos/backend/middleware"
	"github.com/test-tecnico-grupo-lagos/backend/payloads"
)

type APIServer struct {
	listenAddress string
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func NewApiServer(listenAddress string) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
	}
}

func WriteResponse(w http.ResponseWriter, status int, payload any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(payload)
}

// Funcion para usar los modulos que se encargan de recibir las peticiones
func makeHTTPHandlerFunc(handler apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err != nil {
			statusCode := http.StatusBadRequest
			response := &payloads.ResponsePayload{
				StatusCode: statusCode,
				Message:    err.Error(),
				Error:      true,
			}
			WriteResponse(w, statusCode, response)
		}
	}
}

// constantes que seran para definir las rutas del server
const ()

func (server *APIServer) RunServer() {
	router := mux.NewRouter()
	middleware.EnableCORS(router)
	router.Use(middleware.MiddlewareCORS)
	http.ListenAndServe(server.listenAddress, router)
}
