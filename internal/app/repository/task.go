package repository

import (
	"task-management-api/internal/app/model"
	"time"

	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (r *TaskRepository) GetTaskByID(id int) (*model.Task, error) {
	task := new(model.Task)

	err := r.db.
		Model(&model.Task{}).
		Where("id = ?", id).
		First(&task).
		Error
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (r *TaskRepository) ListTasks() ([]model.Task, error) {
	var tasks []model.Task

	err := r.db.
		Model(&model.Task{}).
		Find(&tasks).
		Error
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TaskRepository) CreateTask(taskRequest model.TaskRequest) (*model.Task, error) {
	task := model.Task{
		Title:       taskRequest.Title,
		Description: taskRequest.Description,
		Status:      taskRequest.Status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := r.db.
		Model(&model.Task{}).
		Create(&task).
		Error
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepository) UpdateTask(id int, task model.Task) error {
	err := r.db.
		Model(&model.Task{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"title":       task.Title,
			"description": task.Description,
			"status":      task.Status,
			"updated_at":  time.Now(),
		}).
		Error
	if err != nil {
		return err
	}

	return nil
}

func (r *TaskRepository) UpdateTaskStatus(id int, status model.StatusOfTaskEnum) error {
	err := r.db.
		Model(&model.Task{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":     status,
			"updated_at": time.Now(),
		}).
		Error
	if err != nil {
		return err
	}

	return nil
}

func (r *TaskRepository) DeleteTaskByID(id int) error {
	err := r.db.
		Model(&model.Task{}).
		Where("id = ?", id).
		Delete(&model.Task{}).
		Error
	if err != nil {
		return err
	}

	return nil
}
