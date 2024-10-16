package usecase

import (
	"context"
	domain2 "go-backend-clean-architecture/internal/domain"
	"go-backend-clean-architecture/internal/utils"
	"time"
)

type loginUsecase struct {
	userRepository domain2.UserRepository
	contextTimeout time.Duration
}

func NewLoginUsecase(userRepository domain2.UserRepository, timeout time.Duration) domain2.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (lu *loginUsecase) GetUserByEmail(c context.Context, email string) (domain2.User, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.GetByEmail(ctx, email)
}

func (lu *loginUsecase) CreateAccessToken(user *domain2.User, secret string, expiry int) (accessToken string, err error) {
	return utils.CreateAccessToken(user, secret, expiry)
}

func (lu *loginUsecase) CreateRefreshToken(user *domain2.User, secret string, expiry int) (refreshToken string, err error) {
	return utils.CreateRefreshToken(user, secret, expiry)
}
