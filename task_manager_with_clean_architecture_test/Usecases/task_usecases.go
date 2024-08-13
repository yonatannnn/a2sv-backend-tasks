package Usecases

import "api/task_manager/Domain"

type TaskUsecase struct {
    repo domain.TaskRepository
}

func NewTaskUseCase(repo domain.TaskRepository) *TaskUsecase {
    return &TaskUsecase{repo: repo}
}

func (uc *TaskUsecase) CreateTask(task domain.Task) error {
    return uc.repo.CreateTask(task)
}

func (uc *TaskUsecase) GetTaskByID(id uint) (domain.Task, error) {
    return uc.repo.GetTaskByID(id)
}

func (uc *TaskUsecase) UpdateTask(task domain.Task) error {
    return uc.repo.UpdateTask(task)
}

func (uc *TaskUsecase) DeleteTask(id uint) error {
    return uc.repo.DeleteTask(id)
}

func (uc *TaskUsecase) GetAllTasks() ([]*domain.Task, error) {
    return uc.repo.GetAllTasks()
}
