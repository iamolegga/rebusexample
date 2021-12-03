package repo

import (
	"sync"

	"github.com/iamolegga/rebusexample/internal/domain"
)

type Repo struct {
	m          *sync.RWMutex
	todos      []domain.Todo
	idSequence int
}

func New() *Repo {
	return &Repo{
		&sync.RWMutex{},
		make([]domain.Todo, 0),
		0,
	}
}

func (r *Repo) GetAll() []domain.Todo {
	r.m.RLock()
	defer r.m.RUnlock()

	return r.todos
}

func (r *Repo) GetByID(ID int) *domain.Todo {
	r.m.RLock()
	defer r.m.RUnlock()

	for _, t := range r.todos {
		if t.ID == ID {
			return &t
		}
	}

	return nil
}

func (r *Repo) Create(payload string) domain.Todo {
	r.m.Lock()
	defer r.m.Unlock()

	r.idSequence++
	t := domain.Todo{
		ID:      r.idSequence,
		Payload: payload,
	}
	r.todos = append(r.todos, t)
	return t
}

func (r *Repo) Update(x domain.Todo) *domain.Todo {
	r.m.Lock()
	defer r.m.Unlock()

	for i, t := range r.todos {
		if t.ID == x.ID {
			r.todos[i] = x
			return &x
		}
	}
	return nil
}

func (r *Repo) Delete(ID int) bool {
	r.m.Lock()
	defer r.m.Unlock()

	for i, t := range r.todos {
		if t.ID == ID {
			r.todos = append(r.todos[:i], r.todos[i+1:]...)
			return true
		}
	}

	return false
}
