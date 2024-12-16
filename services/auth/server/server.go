package server

import (
	"auth/db"
	auth "auth/gen"
	"auth/model"
	"auth/utils"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServiceServer struct {
	auth.UnimplementedAuthServiceServer
}

func (s *AuthServiceServer) Register(c context.Context, body *auth.AuthRequest) (*auth.AuthResponse, error) {
	var user = &model.User{
		Name: body.Name,
		Psw:  body.Psw,
	}

	result := db.DB.Create(user)

	if result.RowsAffected == 0 {
		return nil, status.Error(codes.AlreadyExists, "User already exists")
	}

	accToken, refreshToken, err := utils.GenTokens(user.ID)
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "Ошибка создания пользователя")
	}

	return &auth.AuthResponse{
		AccessToken:  accToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthServiceServer) Login(c context.Context, body *auth.AuthRequest) (*auth.AuthResponse, error) {
	var user = &model.User{}

	result := db.DB.Where("name = ?", body.Name).First(&user)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "Пользователь не найден")
	}

	accToken, refreshToken, err := utils.GenTokens(user.ID)
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "Ошибка создания пользователя")
	}

	return &auth.AuthResponse{
		AccessToken:  accToken,
		RefreshToken: refreshToken,
	}, nil

}

func (s *AuthServiceServer) RefreshToken(c context.Context, body *auth.TokenRequest) (*auth.AuthResponse, error) {
	user_id, err := utils.Parse(body.Token)

	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "Ошибка аутентификации")
	}

	accToken, refreshToken, err := utils.GenTokens(uint(user_id))
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "Ошибка создания пользователя")
	}

	return &auth.AuthResponse{
		AccessToken:  accToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthServiceServer) ValidateToken(c context.Context, body *auth.TokenRequest) (*auth.ValidateTokenResponse, error) {
	user_id, err := utils.Parse(body.Token)

	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "Ошибка аутентификации")
	}

	return &auth.ValidateTokenResponse{
		UserId: int32(user_id),
	}, nil
}

func (s *AuthServiceServer) GetUser(c context.Context, body *auth.TokenRequest) (*auth.User, error) {
	user_id, err := utils.Parse(body.Token)

	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "Ошибка аутентификации")
	}

	var user = &model.User{}

	result := db.DB.Where("id = ?", user_id).First(user)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "Пользователь не найден")
	}

	return &auth.User{
		Id:   int32(user.ID),
		Name: user.Name,
	}, nil
}
