package service

import (
	"task-management-api/internal/app/model"
	"task-management-api/internal/app/repository"
	"task-management-api/internal/dbsql"
)

type TaskService struct {
	taskRepository *repository.TaskRepository
}

func NewTaskService() *TaskService {
	db := dbsql.GetDB()
	return &TaskService{
		taskRepository: repository.NewTaskRepository(db),
	}
}

func (s *TaskService) ListTasks() ([]model.Task, error) {
	admins, err := s.taskRepository.ListTasks()
	if err != nil {
		return nil, err
	}
	return admins, nil
}

func (s *TaskService) GetTaskByID(id int) (*model.Task, error) {
	admin, err := s.taskRepository.GetTaskByID(id)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (s *TaskService) CreateTask(taskRequest model.TaskRequest) (*model.Task, error) {
	task, err := s.taskRepository.CreateTask(taskRequest)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *TaskService) UpdateTask(id int, taskUpdateRequest model.TaskUpdateRequest) error {
	task, err := s.taskRepository.GetTaskByID(id)
	if err != nil {
		return err
	}

	if taskUpdateRequest.Title != nil {
		task.Title = *taskUpdateRequest.Title
	}

	if taskUpdateRequest.Description != nil {
		task.Description = *taskUpdateRequest.Description
	}

	if taskUpdateRequest.Status != nil {
		task.Status = *taskUpdateRequest.Status
	}

	err = s.taskRepository.UpdateTask(task.ID, *task)
	if err != nil {
		return err
	}

	return nil
}

func (s *TaskService) UpdateTaskStatus(id int, status model.StatusOfTaskEnum) error {
	task, err := s.taskRepository.GetTaskByID(id)
	if err != nil {
		return err
	}

	err = s.taskRepository.UpdateTaskStatus(task.ID, status)
	if err != nil {
		return err
	}

	return nil
}

func (s *TaskService) DeleteTaskByID(id int) error {
	task, err := s.taskRepository.GetTaskByID(id)
	if err != nil {
		return err
	}

	err = s.taskRepository.DeleteTaskByID(task.ID)
	if err != nil {
		return err
	}

	return nil
}
