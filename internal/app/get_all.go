package app

import "github.com/iamolegga/rebusexample/internal/repo"

type GetAllHandler struct {
	repo *repo.Repo
}

func NewGetAllHandler(repo *repo.Repo) *GetAllHandler {
	return &GetAllHandler{repo: repo}
}

func (h *GetAllHandler) Handle(_ GetAllTodosQuery) (GetAllTodosQueryResult, error) {
	todos := h.repo.GetAll()
	return GetAllTodosQueryResult{todos}, nil
}
