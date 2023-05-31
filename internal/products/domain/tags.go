package products

type TagName string

type Tag struct {
	Name TagName
}

func NewTag(name TagName) Tag {
	return Tag{
		Name: name,
	}
}
