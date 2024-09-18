package request

type CreateUserRequest struct {
	Email    string `validate:"required,min=1,max=100" json:"email"`
	Password string `validate:"required,min=1,max=100" json:"password"`
	Secret   string `validate:"required,min=1,max=100" json:"secret"`
	Token    string `validate:"required,min=1" json:"token"`
	Status   int8   `form:"default=1" json:"status"`
	Role     string `validate:"required,min=1,max=100" json:"role"`
}

type UpdateUserRequest struct {
	ID       uint32 `validate:"required"`
	Email    string `validate:"required,min=1,max=100" json:"email"`
	Password string `validate:"required,min=1,max=100" json:"password"`
	Secret   string `validate:"required,min=1,max=100" json:"secret"`
	Token    string `validate:"required,min=1" json:"token"`
	Status   int8   `json:"status"`
	Role     string `validate:"required,min=1,max=100" json:"role"`
}

type LoginUserRequest struct {
	Email    string `validate:"required,min=1,max=100" json:"email"`
	Password string `validate:"required,min=1,max=100" json:"password"`
}
