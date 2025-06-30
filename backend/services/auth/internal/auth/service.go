package auth

import (
	"errors"
	"gameverse/pkg/configs"
	"gameverse/services/auth/pkg/model"

	"gameverse/pkg/utils"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(input *RegisterRequest) (*model.User, error)
	Login(input *LoginRequest) (string, uuid.UUID, error)
}

type AuthServiceDeps struct {
	UserRepository UserRepository
	Config         *configs.Config
}

type authService struct {
	UserRepository UserRepository
	Config         *configs.Config
}

func NewAuthService(deps *AuthServiceDeps) AuthService {
	return &authService{
		UserRepository: deps.UserRepository,
		Config:         deps.Config,
	}
}

func (s *authService) Register(input *RegisterRequest) (*model.User, error) {
	hashedPwd, err := utils.Hash(input.Password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Email:       input.Email,
		Password:    hashedPwd,
		MobilePhone: input.Phone,
		FirstName:   input.FirsName,
		LastName:    input.LastName,
	}

	err = s.UserRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *authService) Login(input *LoginRequest) (string, uuid.UUID, error) {
	user, err := s.UserRepository.GetByEmail(input.Email)
	if err != nil {
		return "", uuid.Nil, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return "", uuid.Nil, errors.New("invalid credentials")
	}

	token, err := utils.GenerateAccessToken(user.ID, s.Config.Jwt.AccessTTL, s.Config.Jwt.SecretKey)
	if err != nil {
		return "", uuid.Nil, errors.New("failed to generate token")
	}

	return token, user.ID, nil
}
