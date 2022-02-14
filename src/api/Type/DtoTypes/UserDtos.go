package DtoTypes

type CreateUserDto struct {
	Name  string `json:"name,omitempty" validate:"required"`
	Email string `json:"email" validate:"required,email,gte=6"`
	Age   int    `json:"age,omitempty" validate:"required"`
}

type UpdateUserDto struct {
	Name  string `json:"name,omitempty" validate:"required"`
	Email string `json:"email" validate:"required,email,gte=6"`
	Age   int    `json:"age,omitempty" validate:"required"`
}
