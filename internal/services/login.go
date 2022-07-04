package services

import "context"

type LoginInput struct {
	Email    string  `json:"email"`
	Username *string `json:"username"`
}

func (s *Service) Login(ctx context.Context, req *LoginInput) error {
	return nil
}
