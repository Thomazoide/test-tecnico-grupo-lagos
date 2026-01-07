package services

import (
	"github.com/test-tecnico-grupo-lagos/backend/db"
	"github.com/test-tecnico-grupo-lagos/backend/entities"
	"gorm.io/gorm"
)

type CartService struct {
	instance *gorm.DB
}

func NewCartService() *CartService {
	return &CartService{
		instance: db.GetInstance(),
	}
}

func (s *CartService) UpdateBarcodes(cartID uint, barcodes []string) (*entities.Cart, error) {
	var cart entities.Cart
	if err := s.instance.First(&cart, cartID).Error; err != nil {
		return nil, err
	}
	if barcodes == nil {
		barcodes = []string{}
	}
	if err := cart.SetBarcodes(barcodes); err != nil {
		return nil, err
	}
	if err := s.instance.Model(&cart).Update("Barcodes", &cart).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}
