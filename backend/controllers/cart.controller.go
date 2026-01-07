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

type CartController struct {
	service *services.CartService
}

func NewCartController() *CartController {
	return &CartController{
		service: services.NewCartService(),
	}
}

func (c *CartController) HandleUpdateCart(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()
	var barcodes []string
	if err := json.NewDecoder(r.Body).Decode(&barcodes); err != nil {
		return err
	}
	if barcodes == nil {
		barcodes = []string{}
	}
	vars := mux.Vars(r)
	id := vars["id"]
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return err
	}
	cart, err := c.service.UpdateBarcodes(uint(uid), barcodes)
	if err != nil {
		return err
	}
	status := http.StatusOK
	response := &payloads.ResponsePayload{
		StatusCode: status,
		Message:    "Carrito encontrado",
		Data:       cart,
		Error:      false,
	}
	return config.WriteResponse(w, status, response)
}
