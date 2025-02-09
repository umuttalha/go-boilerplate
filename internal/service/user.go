package service

import (
	"errors"
	"go-boilerplate/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) RegisterUser(name, email, password string) (*repository.User, error) {
	// Check if user already exists
	_, err := s.repo.GetUserByEmail(email)
	if err == nil {
		return nil, errors.New("email already registered")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &repository.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Status:   "active",
	}

	if err := s.repo.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) ValidateUser(email, password string) (*repository.User, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
