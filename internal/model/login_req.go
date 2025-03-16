package model

type LoginReq struct {
	Username string `json:"username" validate:"required,min=1,max=256"`
	Password string `json:"password" validate:"required,min=1,max=128"`
}
