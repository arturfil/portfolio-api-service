package project

import (
	"context"
	"time"
)

type Project struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProjectStore interface {
	GetProjects(ctx context.Context) ([]Project, error)
}

type Service struct {
}

// GetProjects implements http.ProjectService
func (Service) GetProjects(ctx context.Context) ([]Project, error) {
	panic("unimplemented")
}

func NewService() Service {
	return Service{}
}
