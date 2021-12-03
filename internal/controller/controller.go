package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/iamolegga/rebusexample/internal/app"
	"github.com/iamolegga/rebusexample/internal/bus"
	"github.com/iamolegga/rebusexample/internal/domain"
)

func New(b bus.Bus) http.Handler {
	var router = mux.NewRouter()

	router.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result, err := b.ExecCreateTodoCommand(app.CreateTodoCommand{Payload: string(body)})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		bytes, _ := json.Marshal(result.Todo)
		_, _ = w.Write(bytes)
	}).Methods(http.MethodPost)

	router.HandleFunc("/todos/{id}", func(w http.ResponseWriter, r *http.Request) {
		ID, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result, err := b.ExecGetTodoQuery(app.GetTodoQuery{ID: ID})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if result.Todo == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		bytes, _ := json.Marshal(result.Todo)
		_, _ = w.Write(bytes)
	}).Methods(http.MethodGet)

	router.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		result, err := b.ExecGetAllTodosQuery(app.GetAllTodosQuery{})
		if err != nil {
		}

		bytes, _ := json.Marshal(result.Todos)
		_, _ = w.Write(bytes)
	}).Methods(http.MethodGet)

	router.HandleFunc("/todos/{id}", func(w http.ResponseWriter, r *http.Request) {
		ID, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		t := domain.Todo{
			ID:      ID,
			Payload: string(body),
		}

		result, err := b.ExecUpdateTodoCommand(app.UpdateTodoCommand{Todo: t})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if result.Todo == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		bytes, _ := json.Marshal(result.Todo)
		_, _ = w.Write(bytes)
	}).Methods(http.MethodPut)

	router.HandleFunc("/todos/{id}", func(w http.ResponseWriter, r *http.Request) {
		ID, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = b.ExecDeleteTodoCommand(app.DeleteTodoCommand{ID: ID})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodDelete)

	return router
}
