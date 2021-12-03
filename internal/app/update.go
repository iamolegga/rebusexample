package app

import "github.com/iamolegga/rebusexample/internal/repo"

type UpdateHandler struct {
	repo *repo.Repo
}

func NewUpdateHandler(repo *repo.Repo) *UpdateHandler {
	return &UpdateHandler{repo: repo}
}

func (h *UpdateHandler) Handle(command UpdateTodoCommand) (UpdateTodoCommandResult, error) {
	t := h.repo.Update(command.Todo)
	return UpdateTodoCommandResult{t}, nil
}
