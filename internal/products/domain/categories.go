package products

type CategoryID int

type Category struct {
	ID             CategoryID
	Name           string
	Description    string
	ParentCategory *Category
}

func NewCategory(ID CategoryID, name string, Description string, parentCategory *Category) Category {
	return Category{
		Name:           name,
		Description:    Description,
		ParentCategory: parentCategory,
	}
}
