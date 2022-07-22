package services

import (
	"context"
	"strings"

	"github.com/google/uuid"

	"go-social-network.com/v1/internal/db"
)

type LoginInput struct {
	Email    string  `json:"email"`
	Username *string `json:"username"`
}

func (in *LoginInput) Prepare() {
	in.Email = strings.ToLower(in.Email)
}

func (s *Service) Login(ctx context.Context, req *LoginInput) (*db.User, error) {
	var out db.User

	req.Prepare()

	exists, err := s.Queries.UserExistsByEmail(ctx, req.Email)
	if err != nil {
		return &out, err
	}

	if exists {
		userByEmail, err := s.Queries.UserByEmail(ctx, req.Email)
		if err != nil {
			return &out, err
		}
		return &userByEmail, nil
	}

	if req.Username == nil {
		return &out, nil
	}

	exists, err = s.Queries.UserExistsByUsername(ctx, *req.Username)
	if err != nil {
		return &out, err
	}

	if exists {
		userByUsername, err := s.Queries.UserByUsername(ctx, *req.Username)
		if err != nil {
			return &out, err
		}

		return &userByUsername, nil
	} else {
		newID, _ := uuid.NewUUID()
		userParams := db.CreateUserParams{
			ID:       newID.String(),
			Email:    req.Email,
			Username: *req.Username,
		}
		createAt, err := s.Queries.CreateUser(ctx, userParams)
		if err != nil {
			return &out, err
		}

		out = db.User{
			ID:        userParams.ID,
			Email:     req.Email,
			Username:  *req.Username,
			CreatedAt: createAt,
			UpdatedAt: createAt,
		}
	}
	return &out, nil
}
