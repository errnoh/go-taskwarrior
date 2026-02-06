// The MIT License (MIT)
// Copyright (C) 2018 Georgy Komarov <jubnzv@gmail.com>

package taskwarrior

import (
	"testing"
)

func TestValidateTask(t *testing.T) {
	// Valid task with all required fields
	validTask := &Task{
		Description: "Test task",
		Status:      "pending",
		Uuid:        "00000000-0000-0000-0000-000000000001",
		Entry:       "20260206T120000Z",
	}

	err := ValidateTask(validTask)
	if err != nil {
		t.Errorf("ValidateTask returned error for valid task: %v", err)
	}

	// Nil task
	err = ValidateTask(nil)
	if err == nil {
		t.Error("ValidateTask should return error for nil task")
	}

	// Empty description
	invalidTask := &Task{
		Status:      "pending",
		Uuid:        "00000000-0000-0000-0000-000000000001",
		Entry:       "20260206T120000Z",
	}

	err = ValidateTask(invalidTask)
	if err == nil {
		t.Error("ValidateTask should return error for empty description")
	}

	// Empty status
	invalidTask = &Task{
		Description: "Test task",
		Uuid:        "00000000-0000-0000-0000-000000000001",
		Entry:       "20260206T120000Z",
	}

	err = ValidateTask(invalidTask)
	if err == nil {
		t.Error("ValidateTask should return error for empty status")
	}

	// Empty uuid
	invalidTask = &Task{
		Description: "Test task",
		Status:      "pending",
		Entry:       "20260206T120000Z",
	}

	err = ValidateTask(invalidTask)
	if err == nil {
		t.Error("ValidateTask should return error for empty uuid")
	}

	// Empty entry
	invalidTask = &Task{
		Description: "Test task",
		Status:      "pending",
		Uuid:        "00000000-0000-0000-0000-000000000001",
	}

	err = ValidateTask(invalidTask)
	if err == nil {
		t.Error("ValidateTask should return error for empty entry")
	}

	// Invalid status
	invalidTask = &Task{
		Description: "Test task",
		Status:      "invalid_status",
		Uuid:        "00000000-0000-0000-0000-000000000001",
		Entry:       "20260206T120000Z",
	}

	err = ValidateTask(invalidTask)
	if err == nil {
		t.Error("ValidateTask should return error for invalid status")
	}
}

func TestValidateTaskRC(t *testing.T) {
	// Valid TaskRC
	validConfig := &TaskRC{
		ConfigPath: "/path/to/taskrc",
		DataLocation: "/path/to/data",
	}

	err := ValidateTaskRC(validConfig)
	if err != nil {
		t.Errorf("ValidateTaskRC returned error for valid config: %v", err)
	}

	// Nil TaskRC
	err = ValidateTaskRC(nil)
	if err == nil {
		t.Error("ValidateTaskRC should return error for nil TaskRC")
	}

	// Empty ConfigPath
	invalidConfig := &TaskRC{
		DataLocation: "/path/to/data",
	}

	err = ValidateTaskRC(invalidConfig)
	if err == nil {
		t.Error("ValidateTaskRC should return error for empty ConfigPath")
	}
}
