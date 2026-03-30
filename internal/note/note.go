package note

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Note struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Tag       string    `json:"tag"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewNote(title, content, tag string) *Note {
	now := time.Now()
	return &Note{
		ID:        uuid.New().String()[:8],
		Title:     title,
		Content:   content,
		Tag:       tag,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}

func (n *Note) String() string {
	return fmt.Sprintf("[%s] %s - %s", n.ID, n.Title, n.Tag)
}