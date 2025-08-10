package dto

type LoginUserDto struct {
	Email    string `json:"email"`
	Password string `json:"password" `
}

type SignUpUserDto struct {
	Username string `json:"username"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" `
}
