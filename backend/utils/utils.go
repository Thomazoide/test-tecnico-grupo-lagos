package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/test-tecnico-grupo-lagos/backend/payloads"
)

type apiFunc func(http.ResponseWriter, *http.Request) error

func WriteResponse(w http.ResponseWriter, status int, payload any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(payload)
}

// Funcion para usar los modulos que se encargan de recibir las peticiones
func MakeHTTPHandlerFunc(handler apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err != nil {
			fmt.Println(err)
			statusCode := http.StatusBadRequest
			response := &payloads.ResponsePayload{
				StatusCode: statusCode,
				Message:    err.Error(),
				Error:      true,
				Data:       err,
			}
			WriteResponse(w, statusCode, response)
		}
	}
}
