package usecase

import (
	"context"
	domain2 "github.com/amitshekhariitbhu/go-backend-clean-architecture/internal/domain"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/utils/tokenutil"
	"time"
)

type signupUsecase struct {
	userRepository domain2.UserRepository
	contextTimeout time.Duration
}

func NewSignupUsecase(userRepository domain2.UserRepository, timeout time.Duration) domain2.SignupUsecase {
	return &signupUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (su *signupUsecase) Create(c context.Context, user *domain2.User) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.Create(ctx, user)
}

func (su *signupUsecase) GetUserByEmail(c context.Context, email string) (domain2.User, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.GetByEmail(ctx, email)
}

func (su *signupUsecase) CreateAccessToken(user *domain2.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (su *signupUsecase) CreateRefreshToken(user *domain2.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
