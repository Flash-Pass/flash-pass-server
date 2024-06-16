package model

type Book struct {
	Base
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedBy   int64  `json:"createdBy"`
}

func NewBook(id uint64, title, content string, createdBy int64) *Book {
	return &Book{
		Base: Base{
			Id: id,
		},
		Title:       title,
		Description: content,
		CreatedBy:   createdBy,
	}
}
