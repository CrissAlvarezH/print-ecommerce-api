package services

import (
	"github.com/CrissAlvarezH/print-ecommerce-api/internal/products/application/ports"
	products "github.com/CrissAlvarezH/print-ecommerce-api/internal/products/domain"
)

type TagService struct {
	repo ports.TagRepository
}

func (s *TagService) List() ([]products.Tag, error) {
	return s.repo.List()
}

func (s *TagService) Create(name products.TagName) (products.Tag, error) {
	return s.repo.Create(name)
}

func (s *TagService) Delete(name products.TagName) error {
	return s.repo.Delete(name)
}
