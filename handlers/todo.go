package handlers

import (
	"log"
	"net/http"
	"strconv"
	"taskManager/data"

	"github.com/gorilla/mux"
)

type Todo struct {
	l *log.Logger
}

func NewTodo(l *log.Logger) *Todo {
	return &Todo{l}
}

func (t *Todo) GetTodo(rw http.ResponseWriter, r *http.Request) {
	t.l.Print("handler GetTodo called")

	tdl := data.GetTodoList()
	err := tdl.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Could not marshal to Json", http.StatusBadRequest)
	}
}

func (t *Todo) PostTodo(rw http.ResponseWriter, r *http.Request) {
	t.l.Print("handler PostTodo called")

	todo := &data.TODO{}
	err := todo.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Could not un-marshal JSON", http.StatusInternalServerError)
	}

	data.AddTodo(todo)
}

func (t *Todo) UpdateTodo(rw http.ResponseWriter, r *http.Request) {
	t.l.Print("handler UpdateTodo called")

	vars := mux.Vars(r)
	id, er := strconv.Atoi(vars["id"])
	if er != nil {
		http.Error(rw, "Unable to convert id to String", http.StatusInternalServerError)
	}

	todo := &data.TODO{}
	err := todo.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Could not un-marshal JSON", http.StatusInternalServerError)
	}

	data.UpdateTodo(id, todo)
}

func (t *Todo) DeleteTodo(rw http.ResponseWriter, r *http.Request) {
	t.l.Print("handler DeleteTodo called")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Could not convert id to String", http.StatusInternalServerError)
	}

	data.DeleteTodo(id)

}
