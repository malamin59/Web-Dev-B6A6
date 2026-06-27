package service

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"spotsync-api/dto"
	"spotsync-api/models"
	"spotsync-api/repository"
	"spotsync-api/utils"
)

type AuthService interface {
	Register(req dto.RegisterRequest) error
	Login(req dto.LoginRequest) (string, error)
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
	fmt.Println("Original:", req.Password)
	fmt.Println("Hashed:", string(hashedPassword))

	// Save user
	if err := s.repo.Create(&user); err != nil {
		return err
	}

	return nil
}

func (s *authService) Login(req dto.LoginRequest) (string, error) {

	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
