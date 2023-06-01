package products

type TextSectionID int64
type TextSectionContentType string

const (
	ContentTypeMarkdown TextSectionContentType = "markdown"
)

type TextSection struct {
	ID          TextSectionID
	Name        string
	Content     string
	ContentType TextSectionContentType
}

func NewTextSection(ID TextSectionID, name string, content string, contentType TextSectionContentType) TextSection {
	return TextSection{
		ID:          ID,
		Name:        name,
		Content:     content,
		ContentType: contentType,
	}
}
