package usecases

import (
	"tr-search-back/internal/domain/user"
)

type DeleteUserUseCase struct {
	userRepository user.Repository
}

func NewDeleteUserUseCase(userRepository user.Repository) *DeleteUserUseCase {
	return &DeleteUserUseCase{userRepository: userRepository}
}

func (uc *DeleteUserUseCase) Execute(id uint) error {
	return uc.userRepository.Delete(id)
}
