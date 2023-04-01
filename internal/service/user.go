package service

import (
	"context"
	v1 "kratos/api/realworld/v1"
	"kratos/internal/errors"
)

func (s *RealWorldService) Login(ctx context.Context, req *v1.LoginRequest) (reply *v1.UserReply, err error) {
	if len(req.User.Email) == 0 {
		return nil, errors.NewHTTPError(422, "email", "cannot be empty")
	}
	rv, err := s.uc.Login(ctx, req.User.Email, req.User.Password)
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Username: rv.UserName,
			Token: rv.Token,
		},
	}, nil
}

func (s *RealWorldService) Register(ctx context.Context, req *v1.RegisterRequest) (reply *v1.UserReply, err error) {
	u, err := s.uc.Register(ctx, req.User.Username, req.User.Email, req.User.Password)
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Email:    u.Email,
			Username: u.UserName,
			Token:    u.Token,
		},
	}, nil
}
func (s *RealWorldService) GetCurrentUser(ctx context.Context, req *v1.GetCurrentUserRequest) (reply *v1.UserReply, err error) {
	return &v1.UserReply{}, nil
}
func (s *RealWorldService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (reply *v1.UserReply, err error) {
	return &v1.UserReply{}, nil
}
