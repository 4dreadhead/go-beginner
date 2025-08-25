package auth

type LoginRequest struct {
	Username string `json:"username" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Username        string `json:"username" validate:"required,email"`
	Password        string `json:"password" validate:"required"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required"`
	FirstName       string `json:"firstName" validate:"required"`
	LastName        string `json:"lastName" validate:"required"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}
