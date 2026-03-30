package calendar

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Time        time.Time `json:"time"`
	Description string    `json:"description"`
	Duration    int       `json:"duration"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewEvent(title, timeStr, description string, duration int) *Event {
	eventTime := time.Now()
	if timeStr != "" {
		if t, err := time.Parse("2006-01-02 15:04", timeStr); err == nil {
			eventTime = t
		}
	}
	return &Event{
		ID:          uuid.New().String()[:8],
		Title:       title,
		Time:        eventTime,
		Description: description,
		Duration:    duration,
		CreatedAt:   time.Now(),
	}
}

func (e *Event) String() string {
	return fmt.Sprintf("%s %s %s (%d min)",
		e.Time.Format("01-02 15:04"),
		e.ID,
		e.Title,
		e.Duration,
	)
}

func (e *Event) EndTime() time.Time {
	return e.Time.Add(time.Duration(e.Duration) * time.Minute)
}