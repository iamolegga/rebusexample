package app

import "github.com/iamolegga/rebusexample/internal/domain"

//GetTodoQuery .
// +rebus:out=../bus
type GetTodoQuery struct {
	ID int
}

type GetTodoQueryResult struct {
	*domain.Todo
}

//GetAllTodosQuery .
// +rebus:out=../bus
type GetAllTodosQuery struct{}

type GetAllTodosQueryResult struct {
	Todos []domain.Todo
}

//CreateTodoCommand .
// +rebus:out=../bus
type CreateTodoCommand struct {
	Payload string
}

type CreateTodoCommandResult struct {
	domain.Todo
}

//UpdateTodoCommand .
// +rebus:out=../bus
type UpdateTodoCommand struct {
	domain.Todo
}

type UpdateTodoCommandResult struct {
	*domain.Todo
}

//DeleteTodoCommand .
// +rebus:out=../bus
type DeleteTodoCommand struct {
	ID int
}
