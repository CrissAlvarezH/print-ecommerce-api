package ports

type ImageFile struct {
	Path        string
	Description string
}

type ImageFilesManager interface {
	SaveAll() ([]ImageFile, error)
	Delete(path string) error
}
