package storage

import (
	"errors"
	"sync"
	"task-dependency-manager/internal/models"
)

type InMemoryStore struct {
	mu    sync.RWMutex
	tasks map[string]*models.Task
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		tasks: make(map[string]*models.Task),
	}
}

func (im *InMemoryStore) GetTask(id string) (*models.Task, error) {
	im.mu.RLock()
	defer im.mu.RUnlock()
	task, ok := im.tasks[id]
	if !ok {
		return &models.Task{}, errors.New("task not found")
	}
	return task, nil
}

func (im *InMemoryStore) CreateTask(task *models.Task) error {
	im.mu.Lock()
	defer im.mu.Unlock()
	im.tasks[task.ID] = task
	return nil
}

func (im *InMemoryStore) UpdateTask(task *models.Task) error {
	im.mu.Lock()
	defer im.mu.Unlock()
	if _, ok := im.tasks[task.ID]; !ok {
		return errors.New("task not found")
	}
	im.tasks[task.ID] = task
	return nil
}

func (im *InMemoryStore) DeleteTask(id string) error {
	im.mu.Lock()
	defer im.mu.Unlock()
	if _, ok := im.tasks[id]; !ok {
		return errors.New("task not found")
	}
	delete(im.tasks, id)
	return nil
}

func (im *InMemoryStore) ListTasks() ([]*models.Task, error) {
	im.mu.RLock()
	defer im.mu.RUnlock()
	tasks := make([]*models.Task, 0, len(im.tasks))
	for _, task := range im.tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}
