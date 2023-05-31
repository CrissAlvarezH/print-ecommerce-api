package services

import (
	"github.com/CrissAlvarezH/print-ecommerce-api/internal/products/application/ports"
	products "github.com/CrissAlvarezH/print-ecommerce-api/internal/products/domain"
	users "github.com/CrissAlvarezH/print-ecommerce-api/internal/users/domain"
	"github.com/Rhymond/go-money"
)

type ProductService struct {
	repo ports.ProductRepository
}

func (s *ProductService) Add(
	sku string, name string, description string, price money.Money,
	discountRate int8, inventoryStatus products.InventoryStatus,
	createdBy users.UserID, variantParentID products.ProductID,
	category products.CategoryName, tags []products.TagName,
	textSections []products.TextSectionID, productAttributes []products.ProductAttributeID,
) (products.Product, error) {
	product, err := s.repo.Add(
		sku, name, description, price, discountRate, inventoryStatus, createdBy,
		variantParentID, category,
	)
	if err != nil {
		return products.Product{}, err
	}

	if err := s.repo.UpdateTags(product.ID, tags); err != nil {
		return products.Product{}, err
	}

	if err := s.repo.UpdateProductAttributes(product.ID, productAttributes); err != nil {
		return products.Product{}, err
	}

	if err := s.repo.UpdateTextSections(product.ID, textSections); err != nil {
		return products.Product{}, err
	}

	product, err = s.repo.GetByID(product.ID)
	if err != nil {
		return products.Product{}, err
	}
	return product, nil
}

func (s *ProductService) GetByID(ID products.ProductID) (products.Product, error) {
	return s.repo.GetByID(ID)
}

func (s *ProductService) Update(
	ID products.ProductID, sku string, name string, description string,
	price money.Money, discountRate int8, inventoryStatus products.InventoryStatus,
	category products.CategoryName, tags []products.TagName, textSections []products.TextSectionID,
	productAttributes []products.ProductAttributeID,
) (products.Product, error) {
	product, err := s.repo.Update(ID, sku, name, description, price, discountRate, inventoryStatus, category)
	if err != nil {
		return products.Product{}, err
	}

	if err := s.repo.UpdateTags(ID, tags); err != nil {
		return products.Product{}, err
	}

	if err := s.repo.UpdateProductAttributes(ID, productAttributes); err != nil {
		return products.Product{}, err
	}

	if err := s.repo.UpdateTextSections(ID, textSections); err != nil {
		return products.Product{}, err
	}

	product, err = s.repo.GetByID(ID)
	if err != nil {
		return products.Product{}, err
	}
	return product, nil
}

func (s *ProductService) AttachImages(ID products.ProductID, images []products.ImageID) error {
	// TODO: continue creating function on this service
	return s.repo.AttachImages(ID, images)
}
