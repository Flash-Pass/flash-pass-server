package book

type Handler struct {
	service BookService
}

type BookService interface{}
