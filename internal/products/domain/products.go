package products

import (
	"errors"
	users "github.com/CrissAlvarezH/print-ecommerce-api/internal/users/domain"
	"github.com/Rhymond/go-money"
	"time"
)

type InventoryStatus string

const (
	InventoryStatusAvailable  InventoryStatus = "AVAILABLE"
	InventoryStatusOutOfStock InventoryStatus = "OUT_OF_STOCK"
	InventoryStatusReserved   InventoryStatus = "RESERVED"
)

type ProductID int64

type Product struct {
	ID                ProductID
	Sku               string
	Name              string
	Description       string
	Price             money.Money
	DiscountRate      int8
	InventoryStatus   InventoryStatus
	CreatedAt         time.Time
	CreatedBy         users.UserID
	VariantParentID   ProductID
	Category          CategoryID
	Tags              []Tag
	TextSections      []TextSection
	ProductAttributes []ProductAttribute
	Images            []ProductImage
	DeletedAt         time.Time
}

func NewProduct(
	ID ProductID, sku string, name string, description string, price money.Money, discountRate int8,
	inventoryStatus InventoryStatus, createdBy users.UserID, variantParentID ProductID, category CategoryID,
	tags []Tag, textSections []TextSection, productAttributes []ProductAttribute, images []ProductImage,
) (Product, error) {
	isLess, err := price.LessThan(money.New(0, money.COP))
	if err != nil {
		return Product{}, err
	}
	if isLess {
		return Product{}, errors.New("price can't be zero")
	}

	return Product{
		ID:                ID,
		Sku:               sku,
		Name:              name,
		Description:       description,
		DiscountRate:      discountRate,
		InventoryStatus:   inventoryStatus,
		CreatedBy:         createdBy,
		VariantParentID:   variantParentID,
		Category:          category,
		CreatedAt:         time.Now(),
		Tags:              tags,
		TextSections:      textSections,
		ProductAttributes: productAttributes,
		Images:            images,
		DeletedAt:         time.Time{},
	}, nil
}

type ProductAttributeID int64

type ProductAttribute struct {
	ID               ProductAttributeID
	AttributeID      AttributeID
	AvailableOptions []AttributeOption
}

func NewProductAttribute(
	ID ProductAttributeID, attributeID AttributeID, availableOptions []AttributeOption,
) ProductAttribute {
	return ProductAttribute{
		ID:               ID,
		AttributeID:      attributeID,
		AvailableOptions: availableOptions,
	}
}
