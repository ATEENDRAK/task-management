package application

import (
	"errors"
	"task-mgmt-service/domain"
)

type taskService struct {
	repo domain.TaskRepository
}

func NewTaskService(repo domain.TaskRepository) domain.TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) CreateTask(title string, status string) (*domain.Task, error) {
	task := &domain.Task{
		Title:  title,
		Status: status,
	}
	err := s.repo.Create(task)
	return task, err
}

func (s *taskService) GetTasks(status string, limit, offset int) ([]domain.Task, error) {
	return s.repo.GetAll(status, limit, offset)
}

func (s *taskService) UpdateTask(id uint, title string, status string) error {
	task, err := s.repo.GetByID(id)
	if err != nil {
		return errors.New("task not found")
	}
	task.Title = title
	task.Status = status
	return s.repo.Update(task)
}

func (s *taskService) DeleteTask(id uint) error {
	return s.repo.Delete(id)
}
