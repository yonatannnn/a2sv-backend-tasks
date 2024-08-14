package Usecases

import (
	domain "api/task_manager/Domain"
	"errors"
)

type UserRepository interface {
	Register(user domain.User) (domain.User, error)
	Login(username, password string) (domain.User, error)
	PromoteUser(userID int) error
	GetUserByID(id int) (domain.User, error)
}


type UserUsecase struct {
	repo UserRepository
}

func NewUserUseCase(repo UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (uc *UserUsecase) GetUserByID(id int) (domain.User, error) {
	return uc.repo.GetUserByID(id)
}

func (uc *UserUsecase) Register(user domain.User) (domain.User, error) {
	if len(user.Username) < 6 {
		return domain.User{} , errors.New("length of username must be greater that 5!")
	}
	if len(user.Password) < 6 {
		return domain.User{} , errors.New("length of password must be greater that 5!")
	}

	return uc.repo.Register(user)
}


func (uc *UserUsecase) Login(username, password string) (domain.User, error) {
	return uc.repo.Login(username , password)
}


func (uc *UserUsecase) PromoteUser(userID int) error {
	return uc.repo.PromoteUser(userID)
}

