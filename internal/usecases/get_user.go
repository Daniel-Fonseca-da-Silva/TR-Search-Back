package usecases

import (
	"tr-search-back/internal/domain/user"
)

type GetUserEmailUseCase struct {
	userRepository user.Repository
}

func NewGetUserEmailUseCase(userRepository user.Repository) *GetUserEmailUseCase {
	return &GetUserEmailUseCase{userRepository: userRepository}
}

func (uc *GetUserEmailUseCase) Execute(email string) (*user.User, error) {
	return uc.userRepository.GetByEmail(email)
}
