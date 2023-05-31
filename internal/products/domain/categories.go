package products

type CategoryName string

type Category struct {
	Name           CategoryName
	Description    string
	ParentCategory CategoryName
}

func NewCategory(name CategoryName, Description string, parentCategory CategoryName) Category {
	return Category{
		Name:           name,
		Description:    Description,
		ParentCategory: parentCategory,
	}
}
