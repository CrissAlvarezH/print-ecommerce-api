package products

type TextSectionID int64

type TextSection struct {
	ID      TextSectionID
	Name    string
	Content string
}

func NewTextSection(ID TextSectionID, name string, content string) TextSection {
	return TextSection{
		ID:      ID,
		Name:    name,
		Content: content,
	}
}
