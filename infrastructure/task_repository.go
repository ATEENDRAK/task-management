package infrastructure

import (
	"fmt"
	"task-mgmt-service/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) domain.TaskRepository {
	return &taskRepository{db: db}
}

func ConnectDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=ateen dbname=task_management port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
	}

	// Auto migrate schema
	err = db.AutoMigrate(&domain.Task{})
	if err != nil {
		fmt.Println("Failed to migrate schema:", err)
	}

	fmt.Println("Connected to PostgresSql database!")
	return db
}

func (r *taskRepository) Create(task *domain.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepository) GetByID(id uint) (*domain.Task, error) {
	var task domain.Task
	err := r.db.First(&task, id).Error
	return &task, err
}

func (r *taskRepository) GetAll(status string, limit, offset int) ([]domain.Task, error) {
	var tasks []domain.Task
	query := r.db
	if status != "" {
		query = query.Where("status = ?", status)
	}
	err := query.Limit(limit).Offset(offset).Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) Update(task *domain.Task) error {
	return r.db.Save(task).Error
}

func (r *taskRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Task{}, id).Error
}
