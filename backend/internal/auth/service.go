package auth

import (
	"errors"
	"gameverse/internal/user"
	"gameverse/pkg/di"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository di.IUserRepository
}

func NewAuthService(userRepository di.IUserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func (s *AuthService) Register(email, name, password string) (string, error) {
	existedUser, _ := s.UserRepository.FindByEmail(email)
	if existedUser != nil {
		return "", errors.New(ErrUserExists)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	newUser := &user.User{
		Email:    email,
		Name:     name,
		Password: string(hashedPassword),
	}

	createdUser, err := s.UserRepository.Create(newUser)
	if err != nil {
		return "", err
	}

	return createdUser.Email, nil
}

func (s *AuthService) Login(email, password string) (string, error) {
	existedUser, err := s.UserRepository.FindByEmail(email)
	if err != nil {
		return "", errors.New(ErrInvalidCredentials)
	}

	err = bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(password))
	if err != nil {
		return "", errors.New(ErrInvalidCredentials)
	}

	return existedUser.Email, nil
}
