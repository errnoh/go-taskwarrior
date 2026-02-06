# AGENTS.md

## Project Overview

This is a Go library providing an API for Taskwarrior (a task management system). The library interacts with Taskwarrior's command-line interface rather than parsing its internal data files directly.

**Module**: `github.com/errnoh/go-taskwarrior`
**Go Version**: 1.25.5

## Essential Commands

### Build
```bash
go build ./...
```

### Test
```bash
go test ./...
go test -v ./...
go test -v -cover ./...
go test -v -race ./...
```

### Get Dependencies
```bash
go get ./...
```

### Run Examples
```bash
go run examples/basics/main.go
```

## Code Organization

### Source Files
- **taskwarrior.go**: Main TaskWarrior struct and core operations (FetchAllTasks, Commit, AddTask, PrintTasks)
- **task.go**: Task struct definition with JSON tags
- **taskrc.go**: TaskRC configuration file parser using reflection
- **task.md**: Taskwarrior JSON format documentation (reference)

### Test Files
- **taskwarrior_test.go**: Tests for TaskWarrior operations
- **taskrc_test.go**: Tests for TaskRC parsing and utilities

### Examples
- **examples/basics/main.go**: Simple example reading user's tasks

### Fixtures
- **fixtures/taskrc/**: Test configuration files (simple_1, err_paths_*, err_permissions_*, redundant_values_*)
- **fixtures/data_1/**: Test data files (completed.data, pending.data)

## Key Patterns and Conventions

### Taskwarrior Interaction Pattern
The library uses Taskwarrior's CLI commands rather than direct file access:
- **Export**: `task rc:<config_path> export` → JSON output → unmarshaled to Task slice
- **Import**: `task import -` → JSON input from stdin → Task slice

### Config File Parsing
- Uses struct tags `taskwarrior:"key"` to map .taskrc file keys to struct fields
- Only parses specific keys that have corresponding struct fields
- Comments (`#`) are stripped from configuration lines
- Include directives are not yet implemented (marked as TODO)

### Task Structure
```go
type Task struct {
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

type Annotation struct {
    Entry        string `json:"entry"`
    Description string `json:"description"`
}
```

All fields are optional except `description`, `status`, `uuid`, and `entry` (required by Taskwarrior JSON format).

### Querying Tasks

`QueryTasks(filter Filter)` retrieves tasks matching the specified filters using Taskwarrior's native filtering capabilities. This is more efficient than fetching all tasks and filtering in memory.

```go
filter := taskwarrior.Filter{
    Project: "work",
    Tags:    []string{"urgent", "high"},
    Status:  "pending",
    UUIDs:   []string{"uuid1", "uuid2"},
}
tasks, err := tw.QueryTasks(filter)
```

**Filter Fields:**
- `Project` - Filter by project name (e.g., "MyProject")
- `Tags` - Filter by tags (e.g., ["urgent", "work"])
- `Status` - Filter by status (e.g., "pending", "completed")
- `UUIDs` - Filter by specific UUIDs (e.g., ["uuid1", "uuid2"])

**Note:** This uses Taskwarrior's command-line filtering for maximum efficiency and compatibility with Taskwarrior's full filter syntax.

### Validation

`ValidateTask()` checks if a Task has all required fields and validates the status value:
- Required: `description`, `status`, `uuid`, `entry`
- Valid status values: `pending`, `completed`, `deleted`, `waiting`, `recurring`
- Returns error if validation fails

`ValidateTaskRC()` checks if TaskRC has valid configuration:
- Required: `ConfigPath`
- Returns error if validation fails

### TaskRC Structure
```go
type TaskRC struct {
    ConfigPath      string // Location of this .taskrc
    DataLocation    string `taskwarrior:"data.location"`
    DependencyTracking string `taskwarrior:"dependency.on"`
    Recall          string `taskwarrior:"recurrence"`
    RecallAfter     string `taskwarrior:"recurrence.limit"`
}
```

Only `DataLocation` is currently supported for parsing.

## Important Gotchas

### Configuration Path Handling
- Config path `~` is expanded to user's home directory via `PathExpandTilda()`
- Empty config path (`""`) defaults to `~/.taskrc`
- Non-existent config paths return errors

### Error Handling
- `FetchAllTasks()` checks for nil TaskWarrior and returns error
- `Commit()` validates JSON marshaling before import
- TaskRC parser validates file existence and permissions

### Deprecated Code
- Uses `io/ioutil` package (deprecated in Go 1.16+)
  - Replace with `os` package functions
  - `ioutil.ReadAll` → `io.ReadAll`
  - `ioutil.ReadFile` → `os.ReadFile`

### Test Requirements
- Tests require Taskwarrior to be installed (`task` command in PATH)
- CI tests install Taskwarrior via apt-get
- Some tests (like `TestParseTaskRC` line 97) have bugs accessing fields before error checking

### Reflection Usage
- `GetAvailableKeys()` uses reflection to discover TaskRC struct fields
- `MapTaskRC()` uses reflection to map config file keys to struct fields
- Field tag `taskwarrior:"key"` is used for mapping

### Taskwarrior JSON Format
- Tasks are single-line JSON objects
- All attribute names are quoted
- Newline characters not permitted in task values
- Unknown fields (UDAs) must be preserved intact

## Development Notes

### Adding New Task Attributes
1. Add field to `Task` struct with appropriate JSON tag
2. Consider `omitempty` for optional fields
3. For User Defined Attributes (UDAs), use the `UDA map[string]interface{}` field which automatically preserves unknown fields

### Adding New TaskRC Options
1. Add field to `TaskRC` struct with `taskwarrior:"key"` tag
2. `GetAvailableKeys()` will automatically discover it
3. `MapTaskRC()` will parse it from config files

### Testing
- Tests use Taskwarrior fixtures from `fixtures/` directory
- Some tests require actual Taskwarrior installation
- Use `-v` flag for verbose test output
- Coverage and race detection available with `-cover` and `-race`

## Known Issues

### Test Failure
`taskrc_test.go:97` has a bug: accesses `result2.ConfigPath` before checking if `err != nil`, causing nil pointer dereference.

### Unimplemented Features
- Include directive parsing in TaskRC files (marked as TODO in `taskrc.go:121`)
- Validation of Taskwarrior command availability

### Deprecated Imports
- `io/ioutil` package used throughout codebase (should migrate to `io` and `os`)
