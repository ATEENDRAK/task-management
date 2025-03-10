package domain

type TaskRepository interface {
	Create(task *Task) error
	GetByID(id uint) (*Task, error)
	GetAll(status string, limit, offset int) ([]Task, error)
	Update(task *Task) error
	Delete(id uint) error
}
