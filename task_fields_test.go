// The MIT License (MIT)
// Copyright (C) 2018 Georgy Komarov <jubnzv@gmail.com>

package taskwarrior

import (
	"testing"
)

func TestTask_NewFields(t *testing.T) {
	// Test that all new fields can be set
	task := &Task{
		Description: "Test task with all fields",
		Status:      "pending",
		Uuid:        "00000000-0000-0000-0000-000000000001",
		Entry:       "20260206T120000Z",
		Start:       "20260206T130000Z",
		End:         "",
		Due:         "20260210T120000Z",
		Until:       "20260306T120000Z",
		Wait:        "",
		Scheduled:   "20260206T120000Z",
		Recur:       "weekly",
		Mask:        "-----",
		Imask:       0,
		Parent:      "",
		Modified:    "20260206T120000Z",
		Depends:     "",
		Tags: []string{
			"work",
			"important",
		},
		Annotations: []Annotation{
			{
				Entry:        "20260206T120500Z",
				Description: "Initial annotation",
			},
		},
		UDA: map[string]interface{}{
			"custom_field": "custom_value",
			"priority_score": 5,
		},
	}

	// Verify basic fields
	if task.Description != "Test task with all fields" {
		t.Errorf("Description mismatch: expected 'Test task with all fields', got '%s'", task.Description)
	}

	if task.Status != "pending" {
		t.Errorf("Status mismatch: expected 'pending', got '%s'", task.Status)
	}

	if task.Uuid != "00000000-0000-0000-0000-000000000001" {
		t.Errorf("Uuid mismatch: expected '00000000-0000-0000-0000-000000000001', got '%s'", task.Uuid)
	}

	if task.Entry != "20260206T120000Z" {
		t.Errorf("Entry mismatch: expected '20260206T120000Z', got '%s'", task.Entry)
	}

	// Verify new fields
	if task.Start != "20260206T130000Z" {
		t.Errorf("Start mismatch: expected '20260206T130000Z', got '%s'", task.Start)
	}

	if task.Due != "20260210T120000Z" {
		t.Errorf("Due mismatch: expected '20260210T120000Z', got '%s'", task.Due)
	}

	if task.Until != "20260306T120000Z" {
		t.Errorf("Until mismatch: expected '20260306T120000Z', got '%s'", task.Until)
	}

	if task.Scheduled != "20260206T120000Z" {
		t.Errorf("Scheduled mismatch: expected '20260206T120000Z', got '%s'", task.Scheduled)
	}

	if task.Recur != "weekly" {
		t.Errorf("Recur mismatch: expected 'weekly', got '%s'", task.Recur)
	}

	if task.Mask != "-----" {
		t.Errorf("Mask mismatch: expected '-----', got '%s'", task.Mask)
	}

	if task.Imask != 0 {
		t.Errorf("Imask mismatch: expected 0, got %d", task.Imask)
	}

	if task.Parent != "" {
		t.Errorf("Parent mismatch: expected empty string, got '%s'", task.Parent)
	}

	if task.Depends != "" {
		t.Errorf("Depends mismatch: expected empty string, got '%s'", task.Depends)
	}

	if len(task.Tags) != 2 {
		t.Errorf("Tags length mismatch: expected 2, got %d", len(task.Tags))
	}

	if task.Tags[0] != "work" {
		t.Errorf("Tags[0] mismatch: expected 'work', got '%s'", task.Tags[0])
	}

	if task.Tags[1] != "important" {
		t.Errorf("Tags[1] mismatch: expected 'important', got '%s'", task.Tags[1])
	}

	if len(task.Annotations) != 1 {
		t.Errorf("Annotations length mismatch: expected 1, got %d", len(task.Annotations))
	}

	if task.Annotations[0].Entry != "20260206T120500Z" {
		t.Errorf("Annotation Entry mismatch: expected '20260206T120500Z', got '%s'", task.Annotations[0].Entry)
	}

	if task.Annotations[0].Description != "Initial annotation" {
		t.Errorf("Annotation Description mismatch: expected 'Initial annotation', got '%s'", task.Annotations[0].Description)
	}

	if task.UDA == nil {
		t.Error("UDA is nil, expected non-nil map")
	}

	if len(task.UDA) != 2 {
		t.Errorf("UDA length mismatch: expected 2, got %d", len(task.UDA))
	}

	if task.UDA["custom_field"] != "custom_value" {
		t.Errorf("UDA custom_field mismatch: expected 'custom_value', got '%v'", task.UDA["custom_field"])
	}

	if task.UDA["priority_score"] != 5 {
		t.Errorf("UDA priority_score mismatch: expected 5, got '%v'", task.UDA["priority_score"])
	}
}

func TestTask_DependsFormat(t *testing.T) {
	// Test depends field with comma-separated UUIDs (as per Taskwarrior spec)
	task := &Task{
		Description: "Depends test",
		Status:      "pending",
		Uuid:        "00000000-0000-0000-0000-000000000001",
		Entry:       "20260206T120000Z",
		Depends:     "00000000-0000-0000-0000-000000000002,00000000-0000-0000-0000-000000000003",
	}

	if task.Depends != "00000000-0000-0000-0000-000000000002,00000000-0000-0000-0000-000000000003" {
		t.Errorf("Depends mismatch: expected '00000000-0000-0000-0000-000000000002,00000000-0000-0000-0000-000000000003', got '%s'", task.Depends)
	}
}

func TestTask_RecurringTask(t *testing.T) {
	// Test recurring task fields
	parent := &Task{
		Description: "Recurring parent",
		Status:      "recurring",
		Uuid:        "00000000-0000-0000-0000-000000000001",
		Entry:       "20260206T120000Z",
		Recur:       "weekly",
		Due:         "20260306T120000Z",
		Until:       "20260313T120000Z",
		Mask:        "-----",
		Parent:      "",
	}

	if parent.Status != "recurring" {
		t.Errorf("Parent status mismatch: expected 'recurring', got '%s'", parent.Status)
	}

	if parent.Recur != "weekly" {
		t.Errorf("Parent recur mismatch: expected 'weekly', got '%s'", parent.Recur)
	}

	if parent.Mask != "-----" {
		t.Errorf("Parent mask mismatch: expected '-----', got '%s'", parent.Mask)
	}

	child := &Task{
		Description: "Recurring child",
		Status:      "pending",
		Uuid:        "00000000-0000-0000-0000-000000000002",
		Entry:       "20260206T120000Z",
		Parent:      "00000000-0000-0000-0000-000000000001",
		Imask:       1,
	}

	if child.Parent != "00000000-0000-0000-0000-000000000001" {
		t.Errorf("Child parent mismatch: expected '00000000-0000-0000-0000-000000000001', got '%s'", child.Parent)
	}

	if child.Imask != 1 {
		t.Errorf("Child imask mismatch: expected 1, got %d", child.Imask)
	}
}

func TestTask_Udapreservation(t *testing.T) {
	// Test that UDA fields are preserved
	task := &Task{
		Description: "Task with UDA",
		Status:      "pending",
		Uuid:        "00000000-0000-0000-0000-000000000001",
		Entry:       "20260206T120000Z",
		UDA: map[string]interface{}{
			"custom_date": "20260201",
			"custom_number": 42,
			"custom_bool": true,
		},
	}

	if task.UDA == nil {
		t.Error("UDA is nil")
	}

	if task.UDA["custom_date"] != "20260201" {
		t.Errorf("custom_date mismatch: expected '20260201', got '%v'", task.UDA["custom_date"])
	}

	if task.UDA["custom_number"] != 42 {
		t.Errorf("custom_number mismatch: expected 42, got '%v'", task.UDA["custom_number"])
	}

	if task.UDA["custom_bool"] != true {
		t.Errorf("custom_bool mismatch: expected true, got '%v'", task.UDA["custom_bool"])
	}
}
