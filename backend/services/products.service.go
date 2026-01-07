package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/test-tecnico-grupo-lagos/backend/payloads"
)

type ProductService struct {
	client *http.Client
}

func NewProductService() *ProductService {
	return &ProductService{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *ProductService) GetProductsFromPage(page int) (*payloads.ProductsPayload, error) {
	if page <= 0 {
		page = 1
	}

	url := fmt.Sprintf("https://world.openfoodfacts.org/api/v2/search?countries=chile&categories=fo&fields=code,product_name,brands,nutriscore_grade,image_url&page=%d", page)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	var out payloads.ProductsPayload
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (s *ProductService) GetProductByCode(code string) (*payloads.ProductByCodePayload, error) {
	url := fmt.Sprintf("https://world.openfoodfacts.org/api/v0/product/%s.json", code)
	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	var out payloads.ProductByCodePayload
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}
