package ports

import (
	products "github.com/CrissAlvarezH/print-ecommerce-api/internal/products/domain"
	users "github.com/CrissAlvarezH/print-ecommerce-api/internal/users/domain"
	"github.com/Rhymond/go-money"
)

type ProductRepository interface {
	List(filters map[string]string, include []string, limit int64, offset int64) ([]products.Product, error)
	GetByID(ID products.ProductID) (products.Product, error)
	Add(
		sku string, name string, description string, price money.Money,
		discountRate int8, inventoryStatus products.InventoryStatus,
		createdBy users.UserID, variantParentID products.ProductID,
		category products.CategoryName,
	) (products.Product, error)
	Update(
		ID products.ProductID, sku string, name string, description string,
		price money.Money, discountRate int8, inventoryStatus products.InventoryStatus,
		category products.CategoryName,
	) (products.Product, error)
	UpdateTags(ID products.ProductID, tags []products.TagName) error
	UpdateTextSections(ID products.ProductID, textSections []products.TextSectionID) error
	UpdateProductAttributes(ID products.ProductID, productAttributes []products.ProductAttributeID) error
	AttachImages(ID products.ProductID, images []products.ImageID) error
	DetachImages(ID products.ProductID, images []products.ImageID) error
	Delete(ID products.ProductID) error
}
