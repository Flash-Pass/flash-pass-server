package model

type Book struct {
	Base
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedBy   uint64 `json:"createdBy"`
}

func NewBook(id uint64, title, content string, createdBy uint64) *Book {
	return &Book{
		Base: Base{
			Id: id,
		},
		Title:       title,
		Description: content,
		CreatedBy:   createdBy,
	}
}
