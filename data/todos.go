package data

import (
	"encoding/json"
	"io"
)

type TODO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (t *TODOs) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(t)
}

func (t *TODO) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(t)
}

func getTodoId() int {
	id := len(todoList) + 1
	return id
}

func AddTodo(t *TODO) {
	t.ID = getTodoId()
	todoList = append(todoList, t)
}

func UpdateTodo(id int, t *TODO) {
	t.ID = id
	todoList[id-1] = t
}

func DeleteTodo(id int) {
	for i, todo := range todoList {
		if todo.ID == id {
			todoList = append(todoList[:i], todoList[i+1:]...)
			break
		}
	}
}

type TODOs []*TODO

func GetTodoList() TODOs {
	return todoList
}

var todoList = []*TODO{
	{
		ID:   1,
		Name: "Complete golang",
	},
	{
		ID:   2,
		Name: "Work on flutter project",
	},
	{
		ID:   3,
		Name: "Read Uniswap V1",
	},
}
