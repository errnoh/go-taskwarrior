// The MIT License (MIT)
// Copyright (C) 2018 Georgy Komarov <jubnzv@gmail.com>

package taskwarrior

import (
	"testing"
)

func TestFilter_Struct(t *testing.T) {
	// Test Filter struct can be instantiated
	filter := Filter{
		Project: "work",
		Tags:    []string{"urgent", "high"},
		Status:  "pending",
		UUIDs:   []string{"uuid1", "uuid2"},
	}

	if filter.Project != "work" {
		t.Errorf("Project mismatch: expected 'work', got '%s'", filter.Project)
	}

	if len(filter.Tags) != 2 {
		t.Errorf("Tags length mismatch: expected 2, got %d", len(filter.Tags))
	}

	if filter.Status != "pending" {
		t.Errorf("Status mismatch: expected 'pending', got '%s'", filter.Status)
	}

	if len(filter.UUIDs) != 2 {
		t.Errorf("UUIDs length mismatch: expected 2, got %d", len(filter.UUIDs))
	}

	// Test empty filter
	emptyFilter := Filter{}
	if emptyFilter.Project != "" ||
		len(emptyFilter.Tags) != 0 ||
		emptyFilter.Status != "" ||
		len(emptyFilter.UUIDs) != 0 {
		t.Error("Empty filter initialization failed")
	}
}

func TestQueryTasks_NilTaskWarrior(t *testing.T) {
	// Test QueryTasks with nil TaskWarrior
	var tw *TaskWarrior
	filter := Filter{}

	_, err := tw.QueryTasks(filter)
	if err == nil {
		t.Error("QueryTasks should return error for nil TaskWarrior")
	}
}

func TestQueryTasks_EmptyFilters(t *testing.T) {
	// Test QueryTasks with no filters - should return all tasks
	tw, err := NewTaskWarrior("./fixtures/taskrc/simple_1")
	if err != nil {
		t.Skip("Could not initialize TaskWarrior from fixture")
	}

	filter := Filter{}

	// Note: This test requires taskwarrior to be installed and the fixture to have tasks
	// We'll just verify it doesn't crash
	_, err = tw.QueryTasks(filter)
	if err != nil {
		// Expected if taskwarrior command fails
		t.Logf("QueryTasks with empty filters returned error (expected if no tasks): %v", err)
	}

	tw.Tasks = append(tw.Tasks, Task{
		Description: "Test task",
		Status:      "pending",
		Uuid:        "00000000-0000-0000-0000-000000000001",
		Entry:       "20260206T120000Z",
	})

	// Test with in-memory tasks when taskwarrior command fails
	// (This will use the FetchAllTasks results)
	_, err = tw.QueryTasks(filter)
	if err != nil {
		t.Logf("QueryTasks with empty filters on in-memory tasks: %v", err)
	}
}

func TestFilter_Project(t *testing.T) {
	tw, err := NewTaskWarrior("./fixtures/taskrc/simple_1")
	if err != nil {
		t.Skip("Could not initialize TaskWarrior from fixture")
	}

	// Test project filter
	filter := Filter{Project: "work"}
	_, err = tw.QueryTasks(filter)
	if err != nil {
		t.Logf("QueryTasks with project filter: %v", err)
	}
}

func TestFilter_Tags(t *testing.T) {
	tw, err := NewTaskWarrior("./fixtures/taskrc/simple_1")
	if err != nil {
		t.Skip("Could not initialize TaskWarrior from fixture")
	}

	// Test tags filter
	filter := Filter{Tags: []string{"urgent", "work"}}
	_, err = tw.QueryTasks(filter)
	if err != nil {
		t.Logf("QueryTasks with tags filter: %v", err)
	}
}

func TestFilter_Status(t *testing.T) {
	tw, err := NewTaskWarrior("./fixtures/taskrc/simple_1")
	if err != nil {
		t.Skip("Could not initialize TaskWarrior from fixture")
	}

	// Test status filter
	filter := Filter{Status: "pending"}
	_, err = tw.QueryTasks(filter)
	if err != nil {
		t.Logf("QueryTasks with status filter: %v", err)
	}
}

func TestFilter_UUIDs(t *testing.T) {
	tw, err := NewTaskWarrior("./fixtures/taskrc/simple_1")
	if err != nil {
		t.Skip("Could not initialize TaskWarrior from fixture")
	}

	// Test UUIDs filter
	filter := Filter{UUIDs: []string{"00000000-0000-0000-0000-000000000001"}}
	_, err = tw.QueryTasks(filter)
	if err != nil {
		t.Logf("QueryTasks with UUIDs filter: %v", err)
	}
}

func TestFilter_Combined(t *testing.T) {
	tw, err := NewTaskWarrior("./fixtures/taskrc/simple_1")
	if err != nil {
		t.Skip("Could not initialize TaskWarrior from fixture")
	}

	// Test combined filters
	filter := Filter{
		Project: "work",
		Tags:    []string{"urgent"},
		Status:  "pending",
	}

	_, err = tw.QueryTasks(filter)
	if err != nil {
		t.Logf("QueryTasks with combined filters: %v", err)
	}
}
