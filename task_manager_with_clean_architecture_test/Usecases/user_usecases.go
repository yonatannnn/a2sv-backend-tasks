package Usecases

import domain "api/task_manager/Domain"




type UserUsecase struct {
	repo domain.UserRepository
}

func NewUserUseCase(repo domain.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (uc *UserUsecase) GetUserByID(id uint) (domain.User, error) {
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

