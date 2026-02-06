# go-taskwarrior

[![GoDoc](https://godoc.org/github.com/errnoh/go-taskwarrior?status.svg)](https://godoc.org/github.com/errnoh/go-taskwarrior)

Golang API for [taskwarrior](https://taskwarrior.org/) database.

## Features

* Full support for Taskwarrior JSON format
* Custom parser for `.taskrc` configuration files
* Read access to taskwarrior database
* Adding/modifying existing tasks
* **Query tasks with filters** (project, tags, status, UUIDs)
* **Field validation** for Task and TaskRC structures
* Support for all Taskwarrior fields:
  - Dates: entry, start, due, until, wait, scheduled, end, modified
  - Recurring tasks: recur, mask, imask, parent
  - Tags and annotations arrays
  - User Defined Attributes (UDAs)
  - Dependency tracking with comma-separated UUIDs
* Comprehensive test suite with fixtures
* Validation helpers:
  - `ValidateTask()` - validates required task fields
  - `ValidateTaskRC()` - validates configuration

## Quickstart

Example program to read the current user's tasks:

```
package main

import (
	"github.com/errnoh/go-taskwarrior"
)

func main() {
	// Initialize new TaskWarrior instance
	tw, err := taskwarrior.NewTaskWarrior("~/.taskrc")
	if err != nil {
		panic(err)
	}

	// Fetch all tasks
	err = tw.FetchAllTasks()
	if err != nil {
		panic(err)
	}

	// Print all tasks
	tw.PrintTasks()
}
```

### Querying Tasks

Filter tasks by project, tags, status, or UUIDs:

```
// Get all pending tasks from the "work" project
filter := taskwarrior.Filter{
    Project: "work",
    Status:  "pending",
}
tasks, err := tw.QueryTasks(filter)

// Get tasks with specific tags
filter := taskwarrior.Filter{
    Tags: []string{"urgent", "high"},
}
tasks, err := tw.QueryTasks(filter)

// Get tasks by UUIDs
filter := taskwarrior.Filter{
    UUIDs: []string{"uuid1", "uuid2"},
}
tasks, err := tw.QueryTasks(filter)
```

### Adding Tasks

To add new task initialize `Task` object with desired values:

```
newTask := &taskwarrior.Task{
    Description: "Buy groceries",
    Project:     "personal",
    Priority:    "H",
    Tags:        []string{"shopping", "weekly"},
    Due:         "20260210T120000Z",
    Urgency:     5.0,
}

tw.AddTask(newTask)
tw.Commit() // Save changes
```

### Validating Tasks

Ensure tasks and configuration are valid before saving:

```
// Validate a task
err := taskwarrior.ValidateTask(newTask)
if err != nil {
    panic(err)
}

// Validate TaskRC configuration
err := taskwarrior.ValidateTaskRC(tw.Config)
if err != nil {
    panic(err)
}
```

### Task Structure

The library supports all Taskwarrior fields:

```
type Task struct {
    Description string  `json:"description"`
    Project     string  `json:"project,omitempty"`
    Status      string  `json:"status,omitempty"`
    Uuid        string  `json:"uuid,omitempty"`
    Urgency     float32 `json:"urgency,omitempty"`
    Priority    string  `json:"priority,omitempty"`
    Start       string  `json:"start,omitempty"`
    Due         string  `json:"due,omitempty"`
    Until       string  `json:"until,omitempty"`
    Wait        string  `json:"wait,omitempty"`
    Scheduled   string  `json:"scheduled,omitempty"`
    Recur       string  `json:"recur,omitempty"`
    Mask        string  `json:"mask,omitempty"`
    Imask       int     `json:"imask,omitempty"`
    Parent      string  `json:"parent,omitempty"`
    Modified    string  `json:"modified,omitempty"`
    Depends     string  `json:"depends,omitempty"`
    Tags        []string `json:"tags,omitempty"`
    Annotations []taskwarrior.Annotation `json:"annotations,omitempty"`
    UDA         map[string]interface{} `json:"-"`
}
```

For more samples see `examples` directory and package tests.
