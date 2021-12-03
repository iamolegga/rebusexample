package app

import "github.com/iamolegga/rebusexample/internal/repo"

type CreateHandler struct {
	repo *repo.Repo
}

func NewCreateHandler(repo *repo.Repo) *CreateHandler {
	return &CreateHandler{repo: repo}
}

func (h *CreateHandler) Handle(command CreateTodoCommand) (CreateTodoCommandResult, error) {
	t := h.repo.Create(command.Payload)
	return CreateTodoCommandResult{t}, nil
}
