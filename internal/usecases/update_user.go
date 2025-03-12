package usecases

import (
	"tr-search-back/internal/domain/user"
)

type UpdateUserUseCase struct {
	userRepository user.Repository
}

func NewUpdateUserUseCase(userRepository user.Repository) *UpdateUserUseCase {
	return &UpdateUserUseCase{userRepository: userRepository}
}

func (uc *UpdateUserUseCase) Execute(user *user.User) error {
	return uc.userRepository.Update(user)
}
