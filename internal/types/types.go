package types

type Student struct {
	Id    int
	Name  string `validate:"required"`//struct tags for validation
	Email string `validate:"required"`//struct tags for validation
	Age   int    `validate:"required"`//struct tags for validation
}
