package repositories

import (
	"tr-search-back/internal/domain/user"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *user.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) Delete(id uint) error {
	result := r.db.Delete(&user.User{}, id)
	return result.Error
}

func (r *UserRepository) GetByEmail(email string) (*user.User, error) {
	var u user.User
	result := r.db.Where("email = ?", email).First(&u)
	if result.Error != nil {
		return nil, result.Error
	}
	return &u, nil
}

func (r *UserRepository) Update(user *user.User) error {
	result := r.db.Save(user)
	return result.Error
}
