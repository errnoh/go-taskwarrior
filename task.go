// The MIT License (MIT)
// Copyright (C) 2018 Georgy Komarov <jubnzv@gmail.com>
//
// Implementation for taskwarrior's Task entries.

package taskwarrior

import (
	"fmt"
)

// Annotation represents a task annotation.
type Annotation struct {
	Entry        string `json:"entry"`
	Description string `json:"description"`
}

// Task representation.
type Task struct {
	Id          int32   `json:"id"`
	Description string  `json:"description"`
	Project     string  `json:"project,omitempty"`
	Status      string  `json:"status,omitempty"`
	Uuid        string  `json:"uuid,omitempty"`
	Urgency     float32 `json:"urgency,omitempty"`
	Priority    string  `json:"priority,omitempty"`
	Due         string  `json:"due,omitempty"`
	Start       string  `json:"start,omitempty"`
	End         string  `json:"end,omitempty"`
	Entry       string  `json:"entry,omitempty"`
	Until       string  `json:"until,omitempty"`
	Wait        string  `json:"wait,omitempty"`
	Scheduled   string  `json:"scheduled,omitempty"`
	Recur       string  `json:"recur,omitempty"`
	Mask        string  `json:"mask,omitempty"`
	Imask       int     `json:"imask,omitempty"`
	Parent      string  `json:"parent,omitempty"`
	Modified    string  `json:"modified,omitempty"`
	Depends     []string `json:"depends,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Annotations []Annotation `json:"annotations,omitempty"`
	UDA         map[string]interface{} `json:"-"`
}

// ValidateTask checks if the task has all required fields.
// Returns error if validation fails.
func ValidateTask(task *Task) error {
	if task == nil {
		return fmt.Errorf("task cannot be nil")
	}

	if task.Description == "" {
		return fmt.Errorf("task description is required")
	}

	if task.Status == "" {
		return fmt.Errorf("task status is required")
	}

	if task.Uuid == "" {
		return fmt.Errorf("task uuid is required")
	}

	if task.Entry == "" {
		return fmt.Errorf("task entry is required")
	}

	// Validate status is one of the allowed values
	validStatuses := map[string]bool{
		"pending":  true,
		"completed": true,
		"deleted":  true,
		"waiting":  true,
		"recurring": true,
	}
	if !validStatuses[task.Status] {
		return fmt.Errorf("invalid task status '%s', must be one of: pending, completed, deleted, waiting, recurring", task.Status)
	}

	return nil
}

// ValidateTaskRC checks if the TaskRC has valid configuration.
// Returns error if validation fails.
func ValidateTaskRC(config *TaskRC) error {
	if config == nil {
		return fmt.Errorf("TaskRC cannot be nil")
	}

	if config.ConfigPath == "" {
		return fmt.Errorf("TaskRC config path is required")
	}

	return nil
}
