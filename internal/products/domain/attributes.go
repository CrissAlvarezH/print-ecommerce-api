package products

type AttributeID int64

type AttributeType string

const (
	AttributeTypeImage  AttributeType = "IMAGE"
	AttributeTypeButton AttributeType = "BUTTON"
	AttributeTypeSelect AttributeType = "SELECT"
)

type Attribute struct {
	ID               AttributeID
	Name             string
	Type             AttributeType
	AttributeOptions []AttributeOption
}

func NewAttribute(ID AttributeID, name string, attributeType AttributeType, options []AttributeOption) Attribute {
	return Attribute{
		ID:               ID,
		Name:             name,
		Type:             attributeType,
		AttributeOptions: options,
	}
}

type AttributeOptionID int64

type AttributeOption struct {
	ID          AttributeOptionID
	Label       string
	Value       string
	ImgPath     string
	AttributeID AttributeID
}

func NewAttributeOption(
	ID AttributeOptionID, label string, value string, imgPath string, attributeID AttributeID,
) AttributeOption {
	return AttributeOption{
		ID:          ID,
		Label:       label,
		Value:       value,
		ImgPath:     imgPath,
		AttributeID: attributeID,
	}
}
