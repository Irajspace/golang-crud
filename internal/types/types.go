package types

type Student struct {
	ID    string `json:"id"`
	Name  string `json:"name" validate:"required"`
	Age   int    `json:"age" validate:"required"`
	Grade string `json:"grade" validate:"required"`
}