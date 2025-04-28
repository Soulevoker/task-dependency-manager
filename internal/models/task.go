package models

import (
	"errors"
	"fmt"
	"strings"
)

type Task struct {
	ID          string `json:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status"`
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
func ValidStatus(status string) bool {
	switch TaskStatus(strings.ToLower(status)) {
	case StatusPending, StatusInProgress, StatusCompleted:
		return true
	default:
		return false
	}
}

func (task *Task) Validate() error {
	if !ValidStatus(task.Status) {
		return errors.New(fmt.Sprintf("invalid task status:%s", task.Status))
	}
	return nil
}
