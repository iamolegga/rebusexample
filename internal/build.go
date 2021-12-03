package internal

import (
	"net/http"

	"github.com/iamolegga/rebusexample/internal/app"
	"github.com/iamolegga/rebusexample/internal/bus"
	"github.com/iamolegga/rebusexample/internal/controller"
	"github.com/iamolegga/rebusexample/internal/repo"
)

func Build() http.Handler {
	r := repo.New()

	createHandler := app.NewCreateHandler(r)
	getHandler := app.NewGetHandler(r)
	getAllHandler := app.NewGetAllHandler(r)
	updateHandler := app.NewUpdateHandler(r)
	deleteHandler := app.NewDeleteHandler(r)

	b := bus.New()
	b.RegisterCreateTodoCommandHandler(createHandler)
	b.RegisterGetTodoQueryHandler(getHandler)
	b.RegisterGetAllTodosQueryHandler(getAllHandler)
	b.RegisterUpdateTodoCommandHandler(updateHandler)
	b.RegisterDeleteTodoCommandHandler(deleteHandler)

	return controller.New(b)
}
