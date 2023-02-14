package usecase

import "github.com/fujitargz/cource-registration-system-backend/internal/domain"

type UserUsecase interface {
	Create(ID string, password string, isAdmin bool) error
	Delete(ID string) error
	FindByID(ID string) (*domain.User, error)
}

type userUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase(userRepository domain.UserRepository) UserUsecase {
	return &userUsecase{userRepository}
}

func (u *userUsecase) Create(ID string, password string, isAdmin bool) error {
	input, err := domain.NewUser(ID, password, isAdmin)
	if err != nil {
		return err
	}
	err = u.userRepository.Save(input.ID, input.PasswordHash, input.IsAdmin)
	if err != nil {
		return err
	}
	return nil
}
func (u *userUsecase) Delete(ID string) error {
	if err := u.userRepository.Delete(ID); err != nil {
		return err
	}
	return nil
}
func (u *userUsecase) FindByID(ID string) (*domain.User, error) {
	user, err := u.userRepository.FindByID(ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
