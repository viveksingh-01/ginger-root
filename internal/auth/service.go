package auth

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Service struct {
	repository Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repository: *repo}
}

func (s *Service) Signup(ctx context.Context, req *SignupRequest) (*User, error) {
	existing, _ := s.repository.FindByEmail(ctx, req.Email)
	if existing != nil {
		return nil, errors.New("email already exists")
	}

	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &User{
		ID:       bson.NewObjectID(),
		Email:    req.Email,
		Name:     req.Name,
		Password: hashedPassword,
		Phone:    req.Phone,
	}

	err = s.repository.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
