package model

type Task struct {
	Base
	PlanId    int64  `json:"planId"`
	BookId    int64  `json:"bookId"`
	Name      string `json:"name"`
	IsActive  bool   `json:"is_active"`
	CreatedBy int64  `json:"created_by"`
}
