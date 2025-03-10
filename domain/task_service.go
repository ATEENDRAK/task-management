package domain

type TaskService interface {
	CreateTask(title string, status string) (*Task, error)
	GetTasks(status string, limit, offset int) ([]Task, error)
	UpdateTask(id uint, title string, status string) error
	DeleteTask(id uint) error
}
