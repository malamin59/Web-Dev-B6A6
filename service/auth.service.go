package service

import (
	"errors"
	"fmt"
	"spotsync-api/dto"
	"spotsync-api/models"
	"spotsync-api/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(req dto.RegisterRequest) error
}

type authService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{
		repo: repo,
	}
}

func (s *authService) Register(req dto.RegisterRequest) error {

	// Check if email already exists
	existingUser, err := s.repo.FindByEmail(req.Email)

	if err == nil && existingUser != nil {
		return errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	// Convert DTO to Model
	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	fmt.Printf("%+v\n", user)

	// Save user
	if err := s.repo.Create(&user); err != nil {
		return err
	}

	return nil
}
