package products

type ImageID int64

type ProductImage struct {
	ID          ImageID
	Path        string
	Description string
	Position    int8
}

func NewProductImage(ID ImageID, path string, description string, position int8) ProductImage {
	return ProductImage{
		ID:          ID,
		Path:        path,
		Description: description,
		Position:    position,
	}
}
