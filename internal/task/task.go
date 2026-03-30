package task

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Priority  string    `json:"priority"`
	DueDate   string    `json:"due_date,omitempty"`
	Tags      []string  `json:"tags,omitempty"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewTask(title, priority, dueDate string, tags []string) *Task {
	now := time.Now()
	return &Task{
		ID:        uuid.New().String()[:8],
		Title:     title,
		Priority:  priority,
		DueDate:   dueDate,
		Tags:      tags,
		Completed: false,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (t *Task) MarkDone() {
	t.Completed = true
	t.UpdatedAt = time.Now()
}

func (t *Task) String() string {
	status := "[ ]"
	if t.Completed {
		status = "[x]"
	}
	priorityIcon := map[string]string{
		"high":   "🔴",
		"medium": "🟡",
		"low":    "🟢",
	}[t.Priority]

	return fmt.Sprintf("%s %s %s %s", status, priorityIcon, t.ID, t.Title)
}