package auth

/***** SignUpDTO *****/
type SignUpDTO struct {
	Login       string `json:"login" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required"`
	Role        string `json:"role" validate:"required"`
	PasswordRaw string `json:"password" validate:"required"`
}

/***** SignInDTO *****/
type SignInDTO struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
