package Usecases

import domain "api/task_manager/Domain"

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
	return uc.repo.Register(user)
}


func (uc *UserUsecase) Login(username, password string) (domain.User, error) {
	return uc.repo.Login(username , password)
}


func (uc *UserUsecase) PromoteUser(userID int) error {
	return uc.repo.PromoteUser(userID)
}

