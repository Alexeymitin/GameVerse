package errmsg

const (
	ErrRefreshTokenRequired  = "refresh token required"
	ErrInvalidOrExpiredToken = "invalid or expired refresh token"
	ErrFailedToGenerateToken = "failed to generate access token"
	ErrFailedToParseTokenTTL = "failed to parse AccessTokenTTL"
	ErrUnauthorized          = "unauthorized"
	ErrForbidden             = "forbidden"
	ErrInternalServer        = "internal server error"
	ErrInvalidCredentials    = "invalid credentials"
	ErrUserNotFound          = "user not found"
	ErrValidationFailed      = "validation failed"
	ErrInvalidRequest        = "invalid request"
)
