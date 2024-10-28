package requests

type RegisterUser struct {
	FirstName *string `json:"FirstName" validate:"omitempty,min=2"`
	LastName  *string `json:"LastName" validate:"omitempty,min=2"`
	UserName  string  `json:"UserName" validate:"min=1,max=32"`
	Password  string  `json:"Password" validate:"min=4,max=32"`
	Email     *string `json:"Email" validate:"omitempty,email"`
	Mobile    *string `json:"Mobile" validate:"omitempty,mobile"`
}

type Login struct {
	UserName string `json:"UserName" validate:"required"`
	Password string `json:"Password" validate:"required"`
}
