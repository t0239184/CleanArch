package repository

import (
	"gorm.io/gorm"

	"github.com/t0239184/CleanArch/app/domain"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindById(id *int64) (*domain.User, error) {
	var user domain.User
	if err := r.db.Take(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindAllUsers() ([]*domain.User, error) {
	var users []*domain.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) CreateUser(user *domain.User) (userId *int64, err error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return &user.Id, nil
}

func (r *UserRepository) UpdateUser(user *domain.User) error {
	if err := r.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) DeleteUser(id *int64) error {
	if err := r.db.Delete(domain.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) UnlockUser(id *int64) error {
	if err := r.db.Model(domain.User{}).Where("id = ?", *id).Update("status", 0).Error; err != nil {
		return err
	}
	return nil
}
