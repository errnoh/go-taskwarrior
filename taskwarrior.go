// The MIT License (MIT)
// Copyright (C) 2018 Georgy Komarov <jubnzv@gmail.com>
//
// Most general definitions to manage list of tasks and taskwarrior instance configuration.
//
// To interact with taskwarrior I decided to use their command-line interface, instead manually parse .data files
// from `data.location` option. This solution looks better because there are few unique .data formats depending of
// taskwarrior version. For more detailed explanations see: https://taskwarrior.org/docs/3rd-party.html.

package taskwarrior

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

// Represents a single taskwarrior instance.
type TaskWarrior struct {
	Config *TaskRC // Configuration options
	Tasks  []Task  // Task JSON entries
}

// Create new empty TaskWarrior instance.
func NewTaskWarrior(configPath string) (*TaskWarrior, error) {
	// Read the configuration file.
	taskRC, err := ParseTaskRC(configPath)
	if err != nil {
		return nil, err
	}

	// Create new TaskWarrior instance.
	tw := &TaskWarrior{Config: taskRC}
	return tw, nil
}

// Fetch all tasks for given TaskWarrior with system `taskwarrior` command call.
func (tw *TaskWarrior) FetchAllTasks() error {
	if tw == nil {
		return fmt.Errorf("Uninitialized taskwarrior database!")
	}

	rcOpt := "rc:" + tw.Config.ConfigPath
	out, err := exec.Command("task", rcOpt, "export").Output()
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(out), &tw.Tasks)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// Pretty print for all tasks represented in given TaskWarrior.
func (tw *TaskWarrior) PrintTasks() {
	out, _ := json.MarshalIndent(tw.Tasks, "", "\t")
	os.Stdout.Write(out)
}

// Add new Task entry to given TaskWarrior.
func (tw *TaskWarrior) AddTask(task *Task) {
	tw.Tasks = append(tw.Tasks, *task)
}

// Save current changes of given TaskWarrior instance.
func (tw *TaskWarrior) Commit() error {
	tasks, err := json.Marshal(tw.Tasks)
	if err != nil {
		return err
	}

	cmd := exec.Command("task", "import", "-")
	cmd.Stdin = bytes.NewBuffer(tasks)
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

// Filter represents query parameters for filtering tasks.
// Used with QueryTasks() to filter exported task data.
type Filter struct {
	Project string // Filter by project (e.g., "MyProject")
	Tags    []string // Filter by tags (e.g., ["urgent", "work"])
	Status  string   // Filter by status (e.g., "pending", "completed")
	UUIDs   []string // Filter by specific UUIDs (e.g., ["uuid1", "uuid2"])
}

// QueryTasks retrieves tasks matching the specified filters using Taskwarrior's
// native filtering capabilities. This is more efficient than fetching all tasks
// and filtering in memory.
//
// Example:
//   filter := taskwarrior.Filter{Project: "work", Tags: ["urgent"]}
//   tasks := tw.QueryTasks(filter)
//
// For more filter syntax, see: https://taskwarrior.org/docs/userguide/#filtering-tasks
func (tw *TaskWarrior) QueryTasks(filter Filter) ([]Task, error) {
	if tw == nil {
		return nil, fmt.Errorf("Uninitialized taskwarrior database!")
	}

	// Build filter string for taskwarrior command
	args := []string{}

	// Add filters before export subcommand
	if filter.Project != "" {
		args = append(args, fmt.Sprintf("project:%s", filter.Project))
	}

	for _, tag := range filter.Tags {
		args = append(args, fmt.Sprintf("+%s", tag))
	}

	if filter.Status != "" {
		args = append(args, fmt.Sprintf("status:%s", filter.Status))
	}

	for _, uuid := range filter.UUIDs {
		args = append(args, fmt.Sprintf("uuid:%s", uuid))
	}

	// Add export subcommand
	args = append(args, "export")

	rcOpt := "rc:" + tw.Config.ConfigPath
	args = append([]string{rcOpt}, args...)

	// Execute task command with filters
	cmd := exec.Command("task", args...)
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	// Unmarshal results
	var tasks []Task
	err = json.Unmarshal([]byte(out), &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
