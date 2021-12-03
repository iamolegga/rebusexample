package app

import "github.com/iamolegga/rebusexample/internal/repo"

type GetHandler struct {
	repo *repo.Repo
}

func NewGetHandler(repo *repo.Repo) *GetHandler {
	return &GetHandler{repo: repo}
}

func (h *GetHandler) Handle(query GetTodoQuery) (GetTodoQueryResult, error) {
	t := h.repo.GetByID(query.ID)
	return GetTodoQueryResult{t}, nil
}
