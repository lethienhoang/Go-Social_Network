package services

import (
	"context"

	"github.com/google/uuid"

	"go-social-network.com/v1/internal/db"
)

type LoginInput struct {
	Email    string  `json:"email"`
	Username *string `json:"username"`
}

func (s *Service) Login(ctx context.Context, req *LoginInput) (*db.Users, error) {
	exists, err := s.Queries.UserExistsByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if exists {
		userByEmail, err := s.Queries.UserByEmail(ctx, req.Email)
		if err != nil {
			return nil, err
		}
		return &userByEmail, nil
	}

	if req.Username != nil {
		exists, err := s.Queries.UserExistsByUsername(ctx, *req.Username)
		if err != nil {
			return nil, err
		}

		if exists {
			userByUsername, err := s.Queries.UserByUsername(ctx, *req.Username)
			if err != nil {
				return nil, err
			}

			return &userByUsername, nil
		} else {
			newID, _ := uuid.NewUUID()
			userParams := db.CreateUserParams{
				ID:       newID.String(),
				Email:    req.Email,
				Username: *req.Username,
			}
			_, err := s.Queries.CreateUser(ctx, userParams)
			if err != nil {
				return nil, err
			}
		}
	}
	return nil, nil
}
