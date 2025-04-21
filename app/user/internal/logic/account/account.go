package account

import (
	"context"
	"proxima-vocabulary/app/user/internal/model/entity"

	"github.com/gogf/gf/v2/os/gtime"
)

func Register(ctx context.Context) (id int, err error) {
	return 1, nil
}

func Login(ctx context.Context) (token string, err error) {
	return "this is a token", nil
}

// get user info
func Info(ctx context.Context, token string) (user *entity.Users, err error) {
	return &entity.Users{
		Id: 1, 
		Username: "admin",
		Password: "123456",
		Email: "admin123456@gmail.com",
		CreatedAt: gtime.New("2025-04-20 14:21:00"),
		UpdatedAt: gtime.New("2025-04-20 14:21:00"),
	}, nil
}