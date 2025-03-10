package usecases

import (
	"tr-search-back/internal/domain/user"
)

type UseCase struct {
	userRepository user.Repository
}

func NewUseCase(userRepository user.Repository) *UseCase {
	return &UseCase{userRepository: userRepository}
}

func (uc *UseCase) Execute(user *user.User) error {
	return uc.userRepository.Create(user)
}
