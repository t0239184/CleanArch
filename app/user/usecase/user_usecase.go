package usecase

import "github.com/t0239184/CleanArch/app/domain"

type UserUsecase struct {
	UserRepository domain.IUserRepository
}

func NewUserUsecase(userRepository domain.IUserRepository) *UserUsecase {
	return &UserUsecase{
		UserRepository: userRepository,
	}
}

func (u *UserUsecase) FindById(id *int64) (*domain.User, error) {
	user, err := u.UserRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserUsecase) FindUsers() ([]*domain.User, error) {
	users, err := u.UserRepository.FindAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserUsecase) CreateUser(user *domain.User) (id *int64, error error) {
	id, err := u.UserRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (u *UserUsecase) UpdateUser(user *domain.User) error {
	existedUser, err := u.UserRepository.FindById(&user.Id)
	if err != nil {
		return err
	}

	if existedUser.Password != user.Password {
		existedUser.Password = user.Password
	}

	err = u.UserRepository.UpdateUser(existedUser)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserUsecase) DeleteUser(id *int64) error {
	existedUser, err := u.UserRepository.FindById(id)
	if err != nil {
		return err
	}

	if err := u.UserRepository.DeleteUser(&existedUser.Id); err != nil {
		return err
	}
	return nil
}

func (u *UserUsecase) UnlockUser(id *int64) error {
	existedUser, err := u.UserRepository.FindById(id)
	if err != nil {
		return err
	}

	if err := u.UserRepository.UnlockUser(&existedUser.Id); err != nil {
		return err
	}
	return nil
}
