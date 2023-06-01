package services

import (
	"github.com/CrissAlvarezH/print-ecommerce-api/internal/products/application/ports"
	products "github.com/CrissAlvarezH/print-ecommerce-api/internal/products/domain"
)

type TextSectionService struct {
	repo ports.TextSectionRepository
}

func (s *TextSectionService) List(limit int64, offset int64) ([]products.TextSection, int64, error) {
	return s.repo.List(limit, offset)
}

func (s *TextSectionService) Create(
	name string, content string, contentType products.TextSectionContentType,
) (products.TextSection, error) {
	return s.repo.Create(name, content, contentType)
}

func (s *TextSectionService) Update(
	ID products.TextSectionID, name string, content string, contentType products.TextSectionContentType,
) (products.TextSection, error) {
	return s.repo.Update(ID, name, content, contentType)
}

func (s *TextSectionService) Delete(ID products.TextSectionID) error {
	return s.repo.Delete(ID)
}
