package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/test-tecnico-grupo-lagos/backend/config"
	"github.com/test-tecnico-grupo-lagos/backend/payloads"
	"github.com/test-tecnico-grupo-lagos/backend/services"
)

type UserController struct {
	service *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		service: services.NewUserService(),
	}
}

func (c *UserController) HandleCreateUser(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()
	var barcodes []string
	if err := json.NewDecoder(r.Body).Decode(&barcodes); err != nil {
		return err
	}
	if barcodes == nil {
		barcodes = []string{}
	}
	user, err := c.service.CreateUser(barcodes)
	if err != nil {
		return err
	}
	status := http.StatusCreated
	response := &payloads.ResponsePayload{
		StatusCode: status,
		Message:    "Usuario creado",
		Data:       user,
		Error:      false,
	}
	return config.WriteResponse(w, status, response)
}

func (c *UserController) HandleFindUserByID(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id := vars["id"]
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return err
	}
	result, err := c.service.GetUserByID(uint(uid))
	if err != nil {
		return err
	}
	status := http.StatusOK
	response := &payloads.ResponsePayload{
		StatusCode: status,
		Message:    "Usuario encontrado",
		Data:       result,
		Error:      false,
	}
	return config.WriteResponse(w, status, response)
}
