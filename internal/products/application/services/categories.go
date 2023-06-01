package services

import (
	"github.com/CrissAlvarezH/print-ecommerce-api/internal/products/application/ports"
	products "github.com/CrissAlvarezH/print-ecommerce-api/internal/products/domain"
)

type CategoryService struct {
	repo ports.CategoryRepository
}

func (s *CategoryService) List() ([]products.Category, error) {
	return s.repo.List()
}

func (s *CategoryService) GetByID(ID products.CategoryID) (products.Category, error) {
	return s.repo.GetByID(ID)
}

func (s *CategoryService) Update(
	ID products.CategoryID, name string, description string, parent products.CategoryID,
) (products.Category, error) {
	return s.repo.Update(ID, name, description, parent)
}

func (s *CategoryService) Delete(ID products.CategoryID) error {
	return s.repo.Delete(ID)
}
