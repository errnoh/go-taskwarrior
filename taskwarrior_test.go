// The MIT License (MIT)
// Copyright (C) 2018 Georgy Komarov <jubnzv@gmail.com>

package taskwarrior

import (
	"encoding/json"
	"os/exec"
	"testing"
)

// Helper that executes `task` with selected config path and return result as new TaskRC instances array.
func UtilTaskCmd(configPath string) ([]Task, error) {
	var out []byte
	if configPath != "" {
		rcOpt := "rc:" + configPath
		out, _ = exec.Command("task", rcOpt, "export").Output()
	} else {
		out, _ = exec.Command("task", "export").Output()
	}

	// Initialize new tasks
	tasks := []Task{}
	err := json.Unmarshal([]byte(out), &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func TestNewTaskWarrior(t *testing.T) {
	config1 := "./fixtures/taskrc/simple_1"
	taskrc1 := &TaskRC{ConfigPath: config1}
	expected1 := &TaskWarrior{Config: taskrc1}
	result1, err := NewTaskWarrior(config1)
	if err != nil {
		t.Errorf("NewTaskWarrior fails with following error: %s", err)
	}
	if expected1.Config.ConfigPath != result1.Config.ConfigPath {
		t.Errorf("Incorrect taskrc path in NewTaskWarrior: expected '%s' got '%s'",
			expected1.Config.ConfigPath, result1.Config.ConfigPath)
	}

	// Incorrect config path
	config2 := "./fixtures/not_exists/33"
	_, err = NewTaskWarrior(config2)
	if err == nil {
		t.Errorf("NewTaskWarrior works with non-existent config '%s'", config2)
	}
}

func TestTaskWarrior_FetchAllTasks(t *testing.T) {
	// Test with fixture configuration
	config1 := "./fixtures/taskrc/simple_1"
	tw1, err := NewTaskWarrior(config1)
	if err != nil {
		t.Errorf("NewTaskWarrior fails with following error: %s", err)
	}

	// Verify TaskRC was parsed correctly
	if tw1.Config.DataLocation != "./fixtures/data_1" {
		t.Errorf("DataLocation mismatch: expected './fixtures/data_1', got '%s'", tw1.Config.DataLocation)
	}

	// Verify we can add tasks with new fields
	task := &Task{
		Description: "Test task",
		Status:      "pending",
		Uuid:        "00000000-0000-0000-0000-000000000001",
		Entry:       "20260206T120000Z",
		Tags: []string{"test", "fixture"},
		UDA: map[string]interface{}{
			"custom_field": "test_value",
		},
	}

	tw1.AddTask(task)

	if len(tw1.Tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tw1.Tasks))
	}

	// Verify the task was added correctly
	if tw1.Tasks[0].Description != "Test task" {
		t.Errorf("Task description mismatch: expected 'Test task', got '%s'", tw1.Tasks[0].Description)
	}

	if len(tw1.Tasks[0].Tags) != 2 {
		t.Errorf("Expected 2 tags, got %d", len(tw1.Tasks[0].Tags))
	}

	if tw1.Tasks[0].UDA["custom_field"] != "test_value" {
		t.Errorf("UDA mismatch: expected 'test_value', got '%v'", tw1.Tasks[0].UDA["custom_field"])
	}

	// Uninitilized database error handling
	tw_buggy, _ := NewTaskWarrior("/tmp/does/not/exists")
	err = tw_buggy.FetchAllTasks()
	if err == nil {
		t.Errorf("Incorrect uninitilized database case handling")
	}
}

func TestTaskWarrior_AddTask(t *testing.T) {
	config1 := "./fixtures/taskrc/simple_1"
	tw1, err := NewTaskWarrior(config1)
	if err != nil {
		t.Errorf("NewTaskWarrior failed: %s", err)
	}

	t1 := &Task{
		Description: "Test task",
		Status:      "pending",
		Uuid:        "00000000-0000-0000-0000-000000000001",
		Entry:       "20260206T120000Z",
	}

	tw1.AddTask(t1)

	if len(tw1.Tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tw1.Tasks))
	}
}
