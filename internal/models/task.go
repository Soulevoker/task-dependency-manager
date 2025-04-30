package models

import (
	"errors"
	"fmt"
	"strings"
)

type Task struct {
	ID          string     `json:"id"`
	Name        string     `json:"name" binding:"required"`
	Description string     `json:"description,omitempty"`
	Status      TaskStatus `json:"status"`
}

// TaskStatus represents the status of a task.
type TaskStatus string

// Valid task status values.
const (
	StatusPending    TaskStatus = "pending"
	StatusInProgress TaskStatus = "in_progress"
	StatusCompleted  TaskStatus = "completed"
)

// ValidStatus checks if a status is valid.
func ValidStatus(status TaskStatus) bool {
	switch status {
	case StatusPending, StatusInProgress, StatusCompleted:
		return true
	default:
		return false
	}
}

func (task *Task) Validate() error {
	if len(strings.TrimSpace(task.Name)) == 0 {
		return errors.New("name is required and cannot be empty")
	}
	if !ValidStatus(task.Status) {
		return fmt.Errorf("invalid task status: %s", task.Status)
	}
	return nil
}
