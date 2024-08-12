package Usecases

import (
	
	domain "api/task_manager/Domain"
)

type TaskRepository interface {
    CreateTask(task domain.Task) error
    GetTaskByID(id int) (domain.Task, error)
    UpdateTask(task domain.Task) error
    DeleteTask(id int) error
    GetAllTasks() ([]*domain.Task, error)
}

type TaskUsecase struct {
    repo TaskRepository
}

func NewTaskUseCase(repo TaskRepository) *TaskUsecase {
    return &TaskUsecase{repo: repo}
}

func (uc *TaskUsecase) CreateTask(task domain.Task) error {
    return uc.repo.CreateTask(task)
}

func (uc *TaskUsecase) GetTaskByID(id int) (domain.Task, error) {
    return uc.repo.GetTaskByID(id)
}

func (uc *TaskUsecase) UpdateTask(task domain.Task) error {
    return uc.repo.UpdateTask(task)
}

func (uc *TaskUsecase) DeleteTask(id int) error {
    return uc.repo.DeleteTask(id)
}

func (uc *TaskUsecase) GetAllTasks() ([]*domain.Task, error) {
    return uc.repo.GetAllTasks()
}
