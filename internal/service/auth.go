package service

import (
	"crypto/rand"
	"encoding/base64"
	"go-boilerplate/internal/repository"
)

func generateToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func (s *Service) CreateUserSession(userID int) (*repository.Session, error) {
	// Delete any existing sessions for this user
	if err := s.repo.DeleteUserSessions(userID); err != nil {
		return nil, err
	}

	token, err := generateToken()
	if err != nil {
		return nil, err
	}

	session := &repository.Session{
		Token:  token,
		UserID: userID,
	}

	if err := s.repo.CreateSession(session); err != nil {
		return nil, err
	}

	return session, nil
}

func (s *Service) ValidateSession(token string) (*repository.Session, error) {
	return s.repo.GetSession(token)
}
