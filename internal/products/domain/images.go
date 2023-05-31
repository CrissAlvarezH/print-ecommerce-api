package products

type ImageID int64

type Image struct {
	ID          ImageID
	Path        string
	Description string
}

func NewImage(ID ImageID, path string, description string) Image {
	return Image{
		ID:          ID,
		Path:        path,
		Description: description,
	}
}
