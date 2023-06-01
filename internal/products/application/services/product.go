package services

import (
	"github.com/CrissAlvarezH/print-ecommerce-api/internal/products/application/ports"
	products "github.com/CrissAlvarezH/print-ecommerce-api/internal/products/domain"
	users "github.com/CrissAlvarezH/print-ecommerce-api/internal/users/domain"
	"github.com/Rhymond/go-money"
	"log"
)

type ProductService struct {
	repo    ports.ProductRepository
	imgRepo ports.ImageRepository
}

func (s *ProductService) List(
	filters map[string]string, include []string, limit int64, offset int64,
) ([]products.Product, int64, error) {
	return s.repo.List(filters, include, limit, offset, true)
}

func (s *ProductService) Add(
	sku string, name string, description string, price money.Money,
	discountRate int8, inventoryStatus products.InventoryStatus,
	createdBy users.UserID, variantParentID products.ProductID,
	category products.CategoryID, tags []products.TagName,
	textSections []products.TextSectionID, productAttributes []products.ProductAttribute,
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

	for _, pa := range productAttributes {
		err := s.repo.AddProductAttribute(product.ID, pa.AttributeID, pa.AvailableOptions)
		if err != nil {
			return products.Product{}, err
		}
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
	category products.CategoryID, tags []products.TagName, textSections []products.TextSectionID,
	productAttributes []products.ProductAttribute,
) (products.Product, error) {
	product, err := s.repo.Update(ID, sku, name, description, price, discountRate, inventoryStatus, category)
	if err != nil {
		return products.Product{}, err
	}

	if err := s.repo.UpdateTags(ID, tags); err != nil {
		return products.Product{}, err
	}

	// Delete all product attributes and insert new ones, with following conditions:
	// Filter product attributes without ID, these will be inserted, otherwise will be updated
	if err := s.repo.DeleteAllProductAttribute(ID); err != nil {
		return products.Product{}, err
	}
	for _, pa := range productAttributes {
		err = nil
		if pa.ID == 0 {
			err = s.repo.AddProductAttribute(product.ID, pa.AttributeID, pa.AvailableOptions)
		} else {
			err = s.repo.UpdateProductAttribute(product.ID, pa.ID, pa.AvailableOptions)
		}
		if err != nil {
			return products.Product{}, err
		}
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

func (s *ProductService) AttachImages(ID products.ProductID, imageManager ports.ImageFilesManager) error {
	imageFiles, err := imageManager.SaveAll()
	if err != nil {
		return err
	}
	imageIDs := make([]products.ImageID, 0, len(imageFiles))
	for _, imgFile := range imageFiles {
		image, err := s.imgRepo.Add(imgFile.Path, imgFile.Description, 0)
		if err != nil {
			log.Println("ERROR: on attach image", "path:", imgFile.Path, "err:", err)
			continue
		}
		imageIDs = append(imageIDs, image.ID)
	}
	return s.repo.AttachImages(ID, imageIDs)
}

func (s *ProductService) DetachImages(
	ID products.ProductID, imageIDs []products.ImageID, imageManager ports.ImageFilesManager,
) error {
	images, err := s.imgRepo.ListByIDs(imageIDs)
	if err != nil {
		return err
	}

	imageFilesDeleted := make([]products.ImageID, 0, len(images))
	for _, img := range images {
		err := imageManager.Delete(img.Path)
		if err != nil {
			return err
		}
		imageFilesDeleted = append(imageFilesDeleted, img.ID)
	}
	err = s.imgRepo.DeleteMany(imageFilesDeleted)
	if err != nil {
		return err
	}

	err = s.repo.DetachImages(ID, imageIDs)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProductService) UpdateImagePositions(
	images []products.ProductImage,
) error {
	return s.imgRepo.UpdatePositions(images)
}

func (s *ProductService) Delete(ID products.ProductID) error {
	return s.repo.MarkAsDelete(ID)
}
