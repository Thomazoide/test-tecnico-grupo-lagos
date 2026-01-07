package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/test-tecnico-grupo-lagos/backend/payloads"
	"github.com/test-tecnico-grupo-lagos/backend/services"
	"github.com/test-tecnico-grupo-lagos/backend/utils"
)

type ProductsController struct {
	service *services.ProductService
}

func NewProductsController() *ProductsController {
	return &ProductsController{
		service: services.NewProductService(),
	}
}

func (c *ProductsController) HandleGetProductsFromPage(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	page := vars["page"]
	pageint, err := strconv.ParseInt(page, 10, 0)
	if err != nil {
		return err
	}
	content, err := c.service.GetProductsFromPage(int(pageint))
	status := http.StatusOK
	response := &payloads.ResponsePayload{
		StatusCode: status,
		Message:    fmt.Sprintf("Productos pagina %d", pageint),
		Data:       content,
		Error:      false,
	}
	return utils.WriteResponse(w, status, response)
}

func (c *ProductsController) HandleGetProductFromCode(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	code := vars["code"]
	content, err := c.service.GetProductByCode(code)
	if err != nil {
		return err
	}
	status := http.StatusOK
	response := &payloads.ResponsePayload{
		StatusCode: status,
		Message:    fmt.Sprintf("Producto %s encontrado", code),
		Data:       content,
		Error:      false,
	}
	return utils.WriteResponse(w, status, response)
}
