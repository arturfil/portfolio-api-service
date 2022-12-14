package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/arturfil/portfolio-api-service/internal/project"
)

type Message struct {
	Message string `json:"message"`
}

type ProjectService interface {
	GetProjects(ctx context.Context) ([]project.Project, error)
}

func (h *Handler) GetProjects(w http.ResponseWriter, r *http.Request) {
	var msg Message
	msg.Message = "Api Works"
	encjson, _ := json.Marshal(msg)
	w.Write(encjson)
}
