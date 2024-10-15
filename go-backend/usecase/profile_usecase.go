package usecase

import (
	"context"
	domain2 "go-backend-clean-architecture/internal/domain"
	"time"
)

type profileUsecase struct {
	userRepository domain2.UserRepository
	contextTimeout time.Duration
}

func NewProfileUsecase(userRepository domain2.UserRepository, timeout time.Duration) domain2.ProfileUsecase {
	return &profileUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (pu *profileUsecase) GetProfileByID(c context.Context, userID string) (*domain2.Profile, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	user, err := pu.userRepository.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &domain2.Profile{Name: user.Name, Email: user.Email}, nil
}
