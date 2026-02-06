// The MIT License (MIT)
// Copyright (C) 2018 Georgy Komarov <jubnzv@gmail.com>
//
// Implementation for taskwarrior's Task entries.

package taskwarrior

// Annotation represents a task annotation.
type Annotation struct {
	Entry        string `json:"entry"`
	Description string `json:"description"`
}

// Task representation.
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
	Depends     string  `json:"depends,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Annotations []Annotation `json:"annotations,omitempty"`
	UDA         map[string]interface{} `json:"-"`
}
