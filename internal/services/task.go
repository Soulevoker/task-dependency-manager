package services

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"slices"
	"strings"
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
		return nil, errors.New("task not found")
	}
	return task, nil
}

func (s *TaskService) CreateTask(task *models.Task) (*models.Task, error) {
	if err := task.Validate(); err != nil {
		return nil, err
	}
	if task.ID == "" {
		task.ID = uuid.New().String()
	}
	task.Status = models.TaskStatus(strings.ToLower(string(task.Status)))
	if err := s.store.CreateTask(task); err != nil {
		return nil, fmt.Errorf("failed to create task: %w", err)
	}
	return task, nil
}

func (s *TaskService) UpdateTask(task *models.Task) (*models.Task, error) {
	if task.ID == "" {
		return nil, errors.New("you must provide a task ID")
	}
	if err := task.Validate(); err != nil {
		return nil, err
	}
	if _, err := s.store.GetTask(task.ID); err != nil {
		return nil, errors.New("task not found")
	}
	task.Status = models.TaskStatus(strings.ToLower(string(task.Status)))
	if err := s.store.UpdateTask(task); err != nil {
		return nil, fmt.Errorf("failed to update task: %w", err)
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

func (s *TaskService) AddDependency(taskID, depID string) error {
	// Verify both tasks exist
	task, err := s.GetTask(taskID)
	if err != nil {
		return fmt.Errorf("task not found: %s", taskID)
	}
	if _, err := s.GetTask(depID); err != nil {
		return fmt.Errorf("dependency task not found: %s", depID)
	}

	// Add dependency
	for _, d := range task.Dependencies {
		if d == depID {
			return nil // Dependency already exists
		}
	}
	task.Dependencies = append(task.Dependencies, depID)

	// Check for cycles
	if s.hasCycle(taskID) {
		// Rollback by removing the dependency
		task.Dependencies = task.Dependencies[:len(task.Dependencies)-1]
		return errors.New("circular dependency detected")
	}

	return s.store.UpdateTask(task)
}

func (s *TaskService) RemoveDependency(taskID, depID string) error {
	task, err := s.GetTask(taskID)
	if err != nil {
		return fmt.Errorf("task not found: %s", taskID)
	}
	if _, err := s.GetTask(depID); err != nil {
		return fmt.Errorf("dependency task not found: %s", depID)
	}

	// Remove dependency
	println(slices.Contains(task.Dependencies, depID))
	for i, d := range task.Dependencies {
		if d == depID {
			task.Dependencies = append(task.Dependencies[:i], task.Dependencies[i+1:]...)
			return s.store.UpdateTask(task)
		}
	}
	return nil // Dependency not found, no action needed
}

// hasCycle checks for cycles in the dependency graph using DFS.
func (s *TaskService) hasCycle(startID string) bool {
	visited := make(map[string]bool)
	recStack := make(map[string]bool)

	var dfs func(id string) bool
	dfs = func(id string) bool {
		visited[id] = true
		recStack[id] = true

		task, err := s.GetTask(id)
		if err != nil {
			return false
		}

		for _, depID := range task.Dependencies {
			if !visited[depID] && dfs(depID) {
				return true
			} else if recStack[depID] {
				return true // Cycle detected
			}
		}

		recStack[id] = false
		return false
	}

	return dfs(startID)
}

func (s *TaskService) SeedTasks() error {
	tasks := []models.Task{
		{
			Name:        "Task 1",
			Description: "Description 1",
			Status:      "pending",
		},
		{
			Name:        "Task 2",
			Description: "Description 2",
			Status:      "in_progress",
		},
	}

	for i := range tasks {
		if _, err := s.CreateTask(&tasks[i]); err != nil {
			return err
		}
	}
	return nil
}
