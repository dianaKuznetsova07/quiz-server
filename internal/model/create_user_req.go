package model

type CreateUserReq struct {
	Username string `json:"username" validate:"required,min=2,max=128"`
	Email    string `json:"email" validate:"required,email"`
	FullName string `json:"full_name" validate:"required,min=1,max=500"`
	Password string `json:"password,omitempty" validate:"required,min=2,max=64"`
}
