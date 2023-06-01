package ports

import (
	products "github.com/CrissAlvarezH/print-ecommerce-api/internal/products/domain"
	users "github.com/CrissAlvarezH/print-ecommerce-api/internal/users/domain"
	"github.com/Rhymond/go-money"
)

type ProductRepository interface {
	List(
		filters map[string]string, include []string, limit int64, offset int64, getTotal bool,
	) ([]products.Product, int64, error)
	GetByID(ID products.ProductID) (products.Product, error)
	Add(
		sku string, name string, description string, price money.Money,
		discountRate int8, inventoryStatus products.InventoryStatus,
		createdBy users.UserID, variantParentID products.ProductID,
		category products.CategoryID,
	) (products.Product, error)
	Update(
		ID products.ProductID, sku string, name string, description string,
		price money.Money, discountRate int8, inventoryStatus products.InventoryStatus,
		category products.CategoryID,
	) (products.Product, error)

	UpdateTags(ID products.ProductID, tags []products.TagName) error
	UpdateTextSections(ID products.ProductID, textSections []products.TextSectionID) error

	AddProductAttribute(
		ID products.ProductID, attribute products.AttributeID, options []products.AttributeOption,
	) error
	UpdateProductAttribute(
		ID products.ProductID, productAttributeID products.ProductAttributeID, options []products.AttributeOption,
	) error
	DeleteAllProductAttribute(ID products.ProductID) error

	AttachImages(ID products.ProductID, images []products.ImageID) error
	DetachImages(ID products.ProductID, images []products.ImageID) error

	MarkAsDelete(ID products.ProductID) error
}

type ImageRepository interface {
	Add(path string, description string, position int8) (products.ProductImage, error)
	ListByIDs(IDs []products.ImageID) ([]products.ProductImage, error)
	DeleteMany(IDs []products.ImageID) error
	UpdatePositions(images []products.ProductImage) error
}

type CategoryRepository interface {
	List() ([]products.Category, error)
	GetByID(ID products.CategoryID) (products.Category, error)
	Update(
		ID products.CategoryID, name string, description string, parent products.CategoryID,
	) (products.Category, error)
	Delete(ID products.CategoryID) error
}

type AttributeRepository interface {
	List() ([]products.Attribute, error)
	GetByID(ID products.AttributeID) (products.Attribute, error)
	Add(
		name string, attributeType products.AttributeType, options []products.AttributeOption,
	) (products.Attribute, error)
	Update(
		ID products.AttributeID, name string, attributeType products.AttributeType,
		options []products.AttributeOption,
	)
}
