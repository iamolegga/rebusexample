package app

import (
	"github.com/iamolegga/rebusexample/internal/repo"
)

type DeleteHandler struct {
	repo *repo.Repo
}

func NewDeleteHandler(repo *repo.Repo) *DeleteHandler {
	return &DeleteHandler{repo: repo}
}

func (h *DeleteHandler) Handle(command DeleteTodoCommand) error {
	_ = h.repo.Delete(command.ID)
	return nil
}
