package services

import (
	"errors"
	"github.com/google/uuid"
	"task-dependency-manager/internal/models"
	"task-dependency-manager/internal/storage"
)

type TaskService struct {
	store *storage.InMemoryStore
}

func NewTaskService(store *storage.InMemoryStore) *TaskService {
	return &TaskService{
		store: store,
	}
}

func (s *TaskService) GetTask(id string) (*models.Task, error) {
	task, err := s.store.GetTask(id)
	if err != nil {
		return &models.Task{}, errors.New("task not found")
	}
	return task, nil
}

func (s *TaskService) CreateTask(task *models.Task) (*models.Task, error) {
	if task.ID == "" {
		task.ID = uuid.New().String()
	}
	if err := s.store.CreateTask(task); err != nil {
		return &models.Task{}, errors.New("task could not be created")
	}
	return task, nil
}

func (s *TaskService) UpdateTask(task *models.Task) (*models.Task, error) {
	if task.ID == "" {
		return &models.Task{}, errors.New("you must provide a task ID")
	}
	if _, err := s.store.GetTask(task.ID); err != nil {
		return &models.Task{}, errors.New("task not found")
	}
	if err := s.store.UpdateTask(task); err != nil {
		return &models.Task{}, errors.New("task could not be updated")
	}
	return task, nil
}

func (s *TaskService) DeleteTask(id string) error {
	if id == "" {
		return errors.New("you must provide a task ID")
	}
	if _, err := s.store.GetTask(id); err != nil {
		return errors.New("task not found")
	}
	if err := s.store.DeleteTask(id); err != nil {
		return errors.New("task could not be deleted")
	}
	return nil
}

func (s *TaskService) ListTasks() ([]*models.Task, error) {
	tasks, err := s.store.ListTasks()
	if err != nil {
		return []*models.Task{}, err
	}
	return tasks, nil
}

func (s *TaskService) SeedTasks() error {
	tasks := make([]models.Task, 0)
	tasks = append(tasks, models.Task{
		Name:        "Task 1",
		Description: "Description 1",
		Status:      "pending",
	})
	tasks = append(tasks, models.Task{
		Name:        "Task 2",
		Description: "Description 2",
		Status:      "in_progress",
	})
	for _, task := range tasks {
		if _, err := s.CreateTask(&task); err != nil {
			return err
		}
	}
	return nil
}
