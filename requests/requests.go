package requests

type RegisterAccount struct {
	FirstName string `json:"FirstName" validate:"required,min=2"`
	LastName  string `json:"LastName" validate:"required,min=2"`
	Age       uint   `json:"Age" validate:"required,teener"`
}
