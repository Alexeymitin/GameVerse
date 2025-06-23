package auth

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Phone    string `json:"mobile_phone" validate:"required,e164"`
	FirsName string `json:"first_name"  validate:"required,min=2,max=50"`
	LastName string `json:"last_name" validate:"required,min=2,max=50"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type RegisterResponse struct {
	AccessToken string `json:"access_token"`
	Message     string `json:"message"`
}
