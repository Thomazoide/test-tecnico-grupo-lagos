package services

import (
	"github.com/google/uuid"
	"github.com/test-tecnico-grupo-lagos/backend/db"
	"github.com/test-tecnico-grupo-lagos/backend/entities"
	"gorm.io/gorm"
)

type UserService struct {
	instance *gorm.DB
}

func NewUserService() *UserService {
	return &UserService{
		instance: db.GetInstance(),
	}
}

func (s *UserService) CreateUser(barcodes []string) (*entities.User, error) {
	var newUser *entities.User
	err := s.instance.Transaction(func(tx *gorm.DB) error {
		user := &entities.User{
			ExternalUUID: uuid.NewString(),
		}
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		cart := &entities.Cart{
			UserID: user.ID,
		}
		if err := cart.SetBarcodes(barcodes); err != nil {
			return err
		}
		if err := tx.Create(cart).Error; err != nil {
			return err
		}
		user.Cart = *cart
		newUser = user
		return nil
	})
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (s *UserService) GetUserByID(id uint) (*entities.User, error) {
	var user entities.User
	if err := s.instance.Preload("Cart").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
